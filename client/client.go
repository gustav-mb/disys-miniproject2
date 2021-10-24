package main

import (
	"bufio"
	pb "chatpb"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var client pb.BroadcastClient
var wait *sync.WaitGroup
var done chan int

func init() {
	wait = &sync.WaitGroup{}
	done = make(chan int)
}

func main() {
	// Parse commandline arguments as '-name <username> -password <password> -server <port>'
	name := flag.String("name", "Anonymous", "The name of the user")
	password := flag.String("password", "admin", "The password for the user")
	port := flag.String("server", pb.Port, "Server port")
	flag.Parse()

	// Create user
	user := newUser(name, password)

	// Connect to Client to Server
	conn, err := grpc.Dial("localhost:"+*port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server. :: %v", err)
	}

	client = pb.NewBroadcastClient(conn)

	pb.ShowLogo()
	connect(user)

	wait.Add(1)
	go startCommandShell(user)
	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
}

// Creates a message stream to the server
func connect(user *pb.User) {
	fmt.Println("Connecting to server...")

	stream, err := client.CreateStream(context.Background(), &pb.Connect{
		User:   user,
		Active: true,
	})

	if err != nil {
		log.Fatalf("Connection failed. :: %v", err)
	}

	fmt.Println("Done.")
	fmt.Println()
	fmt.Println("Type '/help' to show a list of commands.")
	fmt.Println(pb.Line)
	wait.Add(1)
	go receiveMessage(stream)
}

// Mute the client from receiving messages.
func disconnect(user *pb.User) {
	fmt.Println(pb.Line)
	log.Println("Disconnecting from server...")
	var _, err = client.DisconnectStream(context.Background(), user)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("Done.")
	}
	fmt.Println(pb.Line)
}

// Unmute the client from receiving messages.
func reconnect(user *pb.User) {
	fmt.Println(pb.Line)
	log.Println("Connecting to server...")
	var _, err = client.ConnectStream(context.Background(), user)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("Done.")
	}

	fmt.Println(pb.Line)
}

// Mute the client from receiving messages and terminate the process.
func exit(user *pb.User) {
	fmt.Println(pb.Line)
	log.Println("Exitting client...")
	client.DisconnectStream(context.Background(), user)
	log.Println("Done.")
	fmt.Println(pb.Line)
	close(done)
}

// Handle user input and send message, if not a command, to the server.
func startCommandShell(user *pb.User) {
	defer wait.Done()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var input = scanner.Text()

		if strings.HasPrefix(input, "/") {
			switch input {
			case "/exit":
				exit(user)
			case "/connect":
				reconnect(user)
				continue
			case "/disconnect":
				disconnect(user)
				continue
			case "/help":
				pb.Help()
				continue
			default:
				fmt.Println("Unknown command. Type /help to show a list of commands.")
				continue
			}
		} else if len(input) > pb.MaxStrLength {
			fmt.Printf("Error: Exceeding maximum message length of %v characters :: length: %v\n", pb.MaxStrLength, len(input))
			continue
		}

		err := sendMessage(user, input)
		if err != nil {
			break
		}
	}
}

// Publish message to server.
func sendMessage(user *pb.User, content string) error {
	msg := &pb.Message{
		User:      user,
		Content:   content,
		Timestamp: time.Now().String(),
	}

	_, err := client.Publish(context.Background(), msg)

	if err != nil {
		return fmt.Errorf("error sending message. :: %v", err)
	}

	return nil
}

// Receive messages and print them to the console.
func receiveMessage(stream pb.Broadcast_CreateStreamClient) {
	defer wait.Done()

	for {
		msg, err := stream.Recv()

		if err != nil {
			log.Fatalf("Error reading message. :: %v", err)
			break
		}

		log.Printf("%v: %s\n", msg.User.Name, msg.Content)
	}
}

// Create and return a User with an unique id, a specified name and password.
func newUser(name *string, password *string) *pb.User {
	id := sha256.Sum256([]byte(time.Now().String() + *name))
	return &pb.User{
		Id:       hex.EncodeToString(id[:]),
		Name:     *name,
		Password: *password,
	}
}

// NOT IMPLEMENTED YET
// func logMessage() {

// }
