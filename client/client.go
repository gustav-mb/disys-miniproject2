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

var user *pb.User
var client pb.BroadcastClient
var wait *sync.WaitGroup
var done chan int

// Lamport clock
var t int32

func init() {
	wait = &sync.WaitGroup{}
	done = make(chan int)
}

func main() {
	// Parse commandline arguments as '-name <username> -ip <ip address> -server <port>'
	name := flag.String("name", "Anonymous", "The name of the user")
	address := flag.String("ip", "localhost", "The ip address to the server")
	port := flag.String("port", pb.Port, "The port on the ip address")
	flag.Parse()

	// Create user
	user = newUser(name)

	// Print client welcome message
	pb.ShowLogo()
	fmt.Printf("Welcome, %v!\n", user.Name)
	fmt.Println()

	// Connect Client to Server
	conn, err := grpc.Dial(*address+":"+*port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server. :: %v", err)
	}
	defer conn.Close()

	// Create client stub to perform RPCs
	client = pb.NewBroadcastClient(conn)
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

	t++

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

// Disconnect client and end process
func exit(user *pb.User) {
	fmt.Println(pb.Line)
	log.Println("Exitting client...")
	t++
	client.DisconnectStream(context.Background(), &pb.Message{
		User: user,
		Content: "Exit request",
		Timestamp: time.Now().String(),
		Lamport: t,
	})
	
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
			case "/help":
				pb.Help()
				continue
			default:
				fmt.Println("Unknown command. Type /help to show a list of commands.")
				continue
			}
		} else if len(input) > pb.MaxMsgLength {
			fmt.Printf("Error: Exceeding maximum message length of %v characters :: length: %v\n", pb.MaxMsgLength, len(input))
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
		Lamport:   t,
	}

	log.Printf("(%v, Send) Publishing message '%v' to server\n", msg.Lamport, msg.Content)

	msg.Lamport++

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

		if t == 1 {
			log.Printf("(C-Send: %v | S-Broadcast: %v) %v: %s\n", t, msg.Lamport, msg.User.Name, msg.Content)
		} else {
			log.Printf("(S-Broadcast: %v) %v: %s\n", msg.Lamport, msg.User.Name, msg.Content)
		}
		

		t = pb.MaxLamport(msg.Lamport, t) + 1
	}
}

// Create and return a User with an unique id, a specified name and password.
func newUser(name *string) *pb.User {
	id := sha256.Sum256([]byte(time.Now().String() + *name))
	return &pb.User{
		Id:   hex.EncodeToString(id[:]),
		Name: *name,
	}
}
