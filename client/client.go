package main

import (
	"bufio"
	pb "chatpb"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
	utils "utils"

	"google.golang.org/grpc"
)

var client pb.ChittyChatClient
var wait *sync.WaitGroup
var done chan int
var logger *utils.Logger

// Flags
var name = flag.String("name", "Anonymous", "The name of the user")
var address = flag.String("ip", "localhost", "The ip address to the server")
var port = flag.String("port", utils.Port, "The port on the ip address")

// Lamport clock
var t int32

func init() {
	wait = &sync.WaitGroup{}
	done = make(chan int)
}

func main() {
	// Parse commandline arguments as '-name <username> -ip <ip address> -port <port>'
	flag.Parse()

	user := newUser(name)

	// Create logger
	logger = utils.NewLogger("client_" + *name + "_" + user.Id)
	logger.WarningLogger.Println("CLIENT STARTED WITH ID", user.Id)

	// Init close handler to handle signal interrupt
	initCloseHandler(user)

	// Print client welcome message
	utils.ShowLogo()
	fmt.Printf("Welcome, %v!\n", user.Name)
	fmt.Println()

	// Connect Client to Server
	logger.InfoLogger.Printf("Connecting to server with ip address: %v:%v...", *address, *port)
	conn, err := grpc.Dial(*address+":"+*port, grpc.WithInsecure())
	if err != nil {
		logger.ErrorLogger.Fatalf("Could not connect to server. :: %v", err)
	}
	defer conn.Close()

	// Create client stub to perform RPCs and create message stream
	client = pb.NewChittyChatClient(conn)
	connect(user)

	wait.Add(1)
	go startCommandShell(user)
	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	fmt.Println(utils.Line)
	fmt.Println("Exitting Chitty-Chat.")
	fmt.Printf("See you later, %v!\n", user.Name)
	fmt.Println(utils.Line)
	logger.WarningLogger.Println("CLIENT ENDED.")
}

// Creates a message stream to the server
func connect(user *pb.User) {
	fmt.Println("Connecting to server...")
	t++

	logger.InfoLogger.Printf("(%v, Send) Sending connect request to server", t)
	stream, err := client.CreateStream(context.Background(), &pb.ConnectRequest{
		User:    user,
		Lamport: t,
	})

	if err != nil {
		fmt.Println("Connection failed.")
		logger.ErrorLogger.Printf("Connection failed. :: %v", err)
		close(done)
		return
	}
	logger.InfoLogger.Println("Connection established.")
	fmt.Println("Connected successfully.")

	wait.Add(1)
	go receiveMessages(stream)
}

// Disconnect client and end process
func exit(user *pb.User) {
	t++

	logger.InfoLogger.Println("Disconnecting from server...")
	logger.InfoLogger.Printf("(%v, Send) Sending disconnect request to server", t)
	fmt.Println(utils.Line)
	fmt.Println("Disconnecting from server...")

	_, err := client.DisconnectStream(context.Background(), &pb.DisconnectRequest{
		User:    user,
		Lamport: t,
	})

	if err != nil {
		logger.ErrorLogger.Printf("Failed to disconnect. :: %v", err)
		return
	}
	logger.InfoLogger.Println("Disconnected successfully.")

	fmt.Println("Disconnected successfully.")

	close(done)
}

// Handle user input and send message, if not a command, to the server.
func startCommandShell(user *pb.User) {
	defer wait.Done()

	fmt.Println()
	fmt.Println("Type '/help' to show a list of commands.")
	fmt.Println(utils.Line)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var input = scanner.Text()

		if strings.HasPrefix(input, "/") {
			switch input {
			case "/exit":
				exit(user)
				continue
			case "/help":
				utils.Help()
				continue
			case "/info":
				utils.PrintClientInfo(user, t)
				continue
			case "/server":
				utils.PrintServerInfo(*address, *port)
				continue
			default:
				fmt.Println("Unknown command. Type /help to show a list of commands.")
				continue
			}
		} else if len(input) > utils.MaxMsgLength {
			fmt.Printf("Error: Exceeding maximum message length of %v characters :: length: %v\n", utils.MaxMsgLength, len(input))
			continue
		}

		err := sendMessage(user, input)
		if err != nil {
			logger.ErrorLogger.Println(err)
			break
		}
	}
}

// Publish message to server.
func sendMessage(user *pb.User, content string) error {
	logger.InfoLogger.Printf("(%v, Send) Publishing message '%v' to server\n", t, content)

	t++

	msg := newMessage(user, content)
	_, err := client.Publish(context.Background(), msg)

	if err != nil {
		return fmt.Errorf("error sending message. :: %v", err)
	}

	return nil
}

// Receive messages and print them to the console.
func receiveMessages(stream pb.ChittyChat_CreateStreamClient) {
	defer wait.Done()

	for {
		msg, err := stream.Recv()

		if err != nil {
			logger.ErrorPrintf("Error reading messages. :: %v", err)
			close(done)
			break
		}

		if t == 1 {
			logger.InfoLogger.Printf("(C-Send: %v | S-Broadcast: %v) %v: %s\n", t, msg.Lamport, msg.User.Name, msg.Content)
		} else {
			logger.InfoLogger.Printf("(S-Broadcast: %v) %v: %s\n", msg.Lamport, msg.User.Name, msg.Content)
		}

		fmt.Printf("[%v] %v: %s\n", msg.Timestamp, msg.User.Name, msg.Content)

		t = utils.MaxLamport(msg.Lamport, t) + 1
	}
}

// Sets a close handler for abrupt session interruption.
func initCloseHandler(user *pb.User) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		exit(user)
	}()
}

// Creates and returns a User with an unique id and a specified name.
func newUser(name *string) *pb.User {
	id := sha256.Sum256([]byte(time.Now().String() + *name))
	return &pb.User{
		Id:   hex.EncodeToString(id[:]),
		Name: *name,
	}
}

// Creates a new message and returns it.
func newMessage(user *pb.User, content string) *pb.Message {
	return &pb.Message{
		User:      user,
		Content:   content,
		Timestamp: time.Now().Format("02-01-2006 15:04:05"),
		Lamport:   t,
	}
}
