package main

import (
	pb "chatpb"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const (
	Joined string = "Participant %v joined Chitty-Chat at Lamport time L"
	Left   string = "Participant %v left Chitty-Chat at Lamport time L"
)

// Connection struct.
// Contains all information about a client connection.
type Connection struct {
	stream pb.Broadcast_CreateStreamServer
	id     string
	user   *pb.User
	active bool
	error  chan error
}

// Server service struct.
// Contains all connections made.
type Server struct {
	pb.UnimplementedBroadcastServer
	Connections []*Connection
}

func main() {
	// Parse commandline arguments
	port := flag.String("port", pb.Port, "Server port")
	flag.Parse()

	// Init listener
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Could not listen at port %v :: %v", port, err)
	}
	pb.ShowLogo()
	log.Printf("Listening at port: " + *port)

	// Create gRPC server instance
	grpcServer := grpc.NewServer(grpc.StatsHandler(&Handler{}))

	// Create chat server instance
	server := &Server{Connections: make([]*Connection, 0)}

	// Register the server service
	log.Println("Starting ChittyChat server at port " + *port)
	pb.RegisterBroadcastServer(grpcServer, server)

	// gRPC listen and serve
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
}

// Create a connection from Server to Client and save the connection in the server.
func (s *Server) CreateStream(pconn *pb.Connect, stream pb.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		user:   pconn.User,
		error:  make(chan error),
	}

	s.Connections = append(s.Connections, conn)
	s.setStreamActive(conn.user, true)

	return <-conn.error
}

// Set a stream to active or inactive for a user.
// Throws errors if the user doesn't have a registered connection or if trying to set stream to the same activity state it has.
func (s *Server) setStreamActive(user *pb.User, active bool) error {
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
		announcement = Joined
	} else {
		announcement = Left
	}

	s.broadcastServerAnnouncement(s.getConnectionMessage(conn.user, announcement))

	return nil
}

// Set a user's connection to inactive
func (s *Server) DisconnectStream(c context.Context, user *pb.User) (*pb.Close, error) {
	return &pb.Close{}, s.setStreamActive(user, false)
}

// Set a user's connection to active
func (s *Server) ConnectStream(c context.Context, user *pb.User) (*pb.Done, error) {
	return &pb.Done{}, s.setStreamActive(user, true)
}

// Looks up a user connection in the Server service.
// Returns nil if not found.
func (s *Server) findUser(id string) *Connection {
	for i := 0; i < len(s.Connections); i++ {
		if s.Connections[i].id == id {
			return s.Connections[i]
		}
	}

	return nil
}

// Recieve message from client and broadcast it to all clients.
func (s *Server) Publish(ctx context.Context, msg *pb.Message) (*pb.Done, error) {
	log.Printf("Received message '%v' from participant '%v'", msg.Content, msg.User.Name)
	_, err := s.Broadcast(ctx, msg)

	return &pb.Done{}, err
}

// Broadcast a message to all clients on the server.
func (s *Server) Broadcast(ctx context.Context, msg *pb.Message) (*pb.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	// Send message to each client
	for _, conn := range s.Connections {
		wait.Add(1)

		go func(msg *pb.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				log.Printf("Sending message '%v' to user %v", msg.Content, conn.user.Name)

				if err != nil {
					log.Printf("Error with stream %v. Error: %v", conn.stream, err)
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
	return &pb.Close{}, nil
}

// Broadcast a server announcement to all clients.
func (s *Server) broadcastServerAnnouncement(msg string) {
	log.Println(msg)
	s.Broadcast(context.Background(), &pb.Message{
		User:      &pb.User{Id: "S", Name: ">> Server"},
		Content:   msg,
		Timestamp: time.Now().String(),
	})
}

// Formats a connection message string and returns it.
func (s *Server) getConnectionMessage(user *pb.User, announcement string) string {
	return fmt.Sprintf(announcement, user.Name)
}
