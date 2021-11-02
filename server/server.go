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
	"utils"

	"google.golang.org/grpc"
)

var grpcServer *grpc.Server
var service *server
var logger *utils.Logger
var done chan int

// Lamport clock
var t = &utils.Lamport{T: 0}

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
	mu sync.Mutex
}

func main() {
	// Parse commandline arguments
	port := flag.String("port", utils.Port, "The port the server has to run on")
	flag.Parse()

	logger = utils.NewLogger("server")
	logger.WarningLogger.Println("SERVER STARTED.")

	done = make(chan int)

	// Init close handler to handle signal interrupt
	initCloseHandler()

	// Show Chitty-Chat logo
	utils.ShowLogo()

	// Init server
	listener := initServer(port)
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Unknown error. :: %v", err)
		}
	}(listener)

	// Start server
	logger.InfoPrintln("Starting Chitty-Chat server...")
	go startServer(listener)
	logger.InfoPrintln("Chitty-Chat server started.")

	<-done
	logger.WarningLogger.Println("SERVER ENDED.")
}

// Initialize all dependencies for the gRPC server.
func initServer(port *string) net.Listener {
	// Init listener
	logger.InfoPrintln("Listening at port:", *port)
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		logger.ErrorFatalf("Could not listen at port %v. :: %v", *port, err)
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
		logger.ErrorFatalf("Failed to start gRPC server. :: %v", err)
	}
}

// CreateStream Create a connection from Server to a Client and store the connection in the server service.
func (s *server) CreateStream(cr *pb.ConnectRequest, stream pb.ChittyChat_CreateStreamServer) error {
	t.MaxAndIncrement(1)

	logger.InfoPrintf("(%v, Receive) Connection request received from '%v'", t.T, cr.User.Name)

	connection := newConnection(cr.User.Id, stream, cr.User)

	// Add connection to array
	s.addConnection(connection)

	// Set connection to active and broadcast join
	s.setStreamActive(connection.user, true)

	return <-connection.error
}

// DisconnectStream Set a user's connection to inactive
func (s *server) DisconnectStream(_ context.Context, dr *pb.DisconnectRequest) (*pb.Close, error) {
	t.MaxAndIncrement(dr.Lamport)

	logger.InfoPrintf("(%v, Receive) Disconnect request received from '%v'", t.T, dr.User.Name)

	return &pb.Close{}, s.setStreamActive(dr.User, false)
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

	s.broadcastServerAnnouncement(fmt.Sprintf(announcement, conn.user.Name, t.T))

	return nil
}

// Publish Receive message from client and broadcast it to all clients.
func (s *server) Publish(ctx context.Context, msg *pb.Message) (*pb.Done, error) {
	t.MaxAndIncrement(msg.Lamport)	// Receive Publish
	logger.InfoPrintf("(%v, Receive) Received published message '%v' from participant '%v'", t.T, msg.Content, msg.User.Name)
	t.Increment() // Broadcast

	var message = msg
	message.Lamport = t.T
	_, err := s.Broadcast(ctx, message)

	t.Increment() // Send Done back
	logger.InfoLogger.Printf("(%v, Send) Sending back DONE.", t.T)
	return &pb.Done{Lamport: t.T}, err
}

// Broadcast a message to all clients on the server.
func (s *server) Broadcast(_ context.Context, msg *pb.Message) (*pb.Done, error) {
	wait := sync.WaitGroup{}
	doneSending := make(chan int)

	// Send message to each client connection
	for _, conn := range s.connections {
		wait.Add(1)

		go func(msg *pb.Message, conn *connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				logger.InfoPrintf("(%v, Send) Broadcasting message '%v' from '%v' to '%v'", msg.Lamport, msg.Content, msg.User.Name, conn.user.Name)

				if err != nil {
					logger.ErrorPrintf("Error with stream %v. :: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)
	}

	go func() {
		wait.Wait()
		close(doneSending)
	}()

	<-doneSending
	return &pb.Done{}, nil
}

// Broadcast a server announcement to all clients.
func (s *server) broadcastServerAnnouncement(msg string) {
	t.Increment()
	s.Broadcast(context.Background(), newMessage(msg))
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
	logger.WarningPrintln("Shutting down server in 3 seconds...")
	time.Sleep(3 * time.Second)
	grpcServer.Stop()
	logger.WarningPrintln("Server shutdown successfully.")
	fmt.Println(utils.Line)
	close(done)
}

// Adds a connection to the connections array and sorts it by id.
func (s *server) addConnection(c *connection) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.connections = append(s.connections, c)
	sort.SliceStable(s.connections, func(i, j int) bool {
		return s.connections[i].id < s.connections[j].id
	})
}

// Sets a close handler for abrupt session interruption.
func initCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		service.shutdown()
	}()
}

// Create a new connection.
func newConnection(id string, stream pb.ChittyChat_CreateStreamServer, user *pb.User) *connection {
	return &connection{
		id:     id,
		stream: stream,
		user:   user,
		active: false,
		error:  make(chan error),
	}
}

// Create a new server message.
func newMessage(content string) *pb.Message {
	return &pb.Message{
		User:      &pb.User{Id: "S", Name: ">> Server"},
		Content:   content,
		Timestamp: time.Now().Format("02-01-2006 15:04:05"),
		Lamport:   t.T,
	}
}
