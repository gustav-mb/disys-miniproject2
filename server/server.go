package main

import (
	pb "chatpb"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"
	utils "utils"

	"google.golang.org/grpc"
)

var grpcServer *grpc.Server
var service *server

// Lamport clock
var t int32

// Server announcements
const (
	joined string = "Participant %v joined Chitty-Chat at lamport time %v (S-Receive)"
	left   string = "Participant %v left Chitty-Chat at lamport time %v (S-Receive)"
)

// Connection struct.
// Contains all information about a client connection.
type connection struct {
	id     string
	stream pb.ChittyChat_CreateStreamServer
	user   *pb.User
	active bool
	error  chan error
}

// Server service struct.
// Contains all connections made.
type server struct {
	pb.UnimplementedChittyChatServer
	connections []*connection
}

func main() {
	// Parse commandline arguments
	port := flag.String("port", utils.Port, "The port the server has to run on")
	flag.Parse()

	done := make(chan int)

	// Init close handler to handle signal interrupt
	initCloseHandler()

	// Show Chitty-Chat logo
	utils.ShowLogo()

	// Init server
	listener := initServer(port)
	defer listener.Close()

	// Start server
	log.Println("Starting Chitty-Chat server...")
	go startServer(listener)
	log.Println("Done.")

	<-done
	service.shutdown()
}

// Initialize all dependencies for the gRPC server.
func initServer(port *string) net.Listener {
	// Init listener
	log.Printf("Listening at port: " + *port)
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Could not listen at port %v :: %v", port, err)
	}

	// Create gRPC server instance
	grpcServer = grpc.NewServer()

	// Create chat service instance
	service = &server{connections: make([]*connection, 0)}

	// Register the server service on the grpc server
	pb.RegisterChittyChatServer(grpcServer, service)

	return listener
}

// Start the gRPC server.
func startServer(listener net.Listener) {
	err := grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to start gRPC server. :: %v", err)
	}
}

// Create a connection from Server to a Client and store the connection in the server service.
func (s *server) CreateStream(cr *pb.ConnectRequest, stream pb.ChittyChat_CreateStreamServer) error {
	// Adjust lamport clock
	t = utils.MaxLamport(t, 1) + 1
	log.Printf("(%v, Receive) Connection request received from '%v'", t, cr.User.Name)

	connection := newConnection(cr.User.Id, stream, cr.User)

	// Append connection to array and sort connections array by id
	s.connections = append(s.connections, connection)
	sort.SliceStable(s.connections, func(i, j int) bool {
		return s.connections[i].id < s.connections[j].id
	})

	// Set connection to active and broadcast join
	s.setStreamActive(connection.user, true)

	return <-connection.error
}

// Set a connection to active (true) or inactive (false) for a client.
// Throws errors if the user doesn't have a registered connection or if trying to set stream to the same activity state it has.
func (s *server) setStreamActive(user *pb.User, active bool) error {
	var conn = s.findUser(user.Id)
	if conn == nil {
		return fmt.Errorf("user connection could not be found :: user: %v (id: %v)", user.Name, user.Id)
	}

	if conn.active && active {
		return fmt.Errorf("already connected to server")
	}

	if !conn.active && !active {
		return fmt.Errorf("not connected to any server")
	}

	conn.active = active
	var announcement string
	if active {
		announcement = joined
	} else {
		announcement = left
	}

	s.broadcastServerAnnouncement(fmt.Sprintf(announcement, conn.user.Name, t))

	return nil
}

// Set a user's connection to inactive
func (s *server) DisconnectStream(ctx context.Context, dr *pb.DisconnectRequest) (*pb.Close, error) {
	t = utils.MaxLamport(t, dr.Lamport) + 1
	log.Printf("(%v, Receive) Disconnect request received from '%v'", t, dr.User.Name)

	return &pb.Close{}, s.setStreamActive(dr.User, false)
}

// Recieve message from client and broadcast it to all clients.
func (s *server) Publish(ctx context.Context, msg *pb.Message) (*pb.Done, error) {
	log.Printf("(%v, Receive), Received message '%v' from participant '%v'", msg.Lamport, msg.Content, msg.User.Name)

	t = utils.MaxLamport(t, msg.Lamport) + 1
	t++

	var message = msg
	message.Lamport = t
	_, err := s.Broadcast(ctx, message)

	return &pb.Done{}, err
}

// Broadcast a message to all clients on the server.
func (s *server) Broadcast(ctx context.Context, msg *pb.Message) (*pb.Done, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	// Send message to each client connection
	for _, conn := range s.connections {
		wait.Add(1)

		go func(msg *pb.Message, conn *connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				log.Printf("(%v, Send) Broadcasting message '%v' to '%v'", msg.Lamport, msg.Content, conn.user.Name)

				if err != nil {
					log.Printf("Error with stream %v. :: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &pb.Done{}, nil
}

// Broadcast a server announcement to all clients.
func (s *server) broadcastServerAnnouncement(msg string) {
	t++

	s.Broadcast(context.Background(), &pb.Message{
		User:      &pb.User{Id: "S", Name: ">> Server"},
		Content:   msg,
		Timestamp: time.Now().String(),
		Lamport:   t,
	})
}

// Looks up and returns a connection equal to the provided id in the Server service using binary search.
// Returns nil if not found.
func (s *server) findUser(id string) *connection {
	index := sort.Search(len(s.connections), func(i int) bool {
		return s.connections[i].id >= id
	})

	if index < len(s.connections) && s.connections[index].id == id {
		return s.connections[index]
	}

	return nil
}

// Shuts down the server after 3 seconds
func (s *server) shutdown() {
	s.broadcastServerAnnouncement("WARNING: Server shutting down in 3 seconds! All users will be disconnected.")
	time.Sleep(3 * time.Second)
	fmt.Println(utils.Line)
	log.Println("Shutting down server...")
	grpcServer.Stop()
	log.Println("Done.")
	fmt.Println(utils.Line)
}

// Sets a close handler for abrupt session interruption.
func initCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		service.shutdown()
		os.Exit(0)
	}()
}

// Create new connection
func newConnection(id string, stream pb.ChittyChat_CreateStreamServer, user *pb.User) *connection {
	return &connection{
		id:     id,
		stream: stream,
		user:   user,
		active: false,
		error:  make(chan error),
	}
}
