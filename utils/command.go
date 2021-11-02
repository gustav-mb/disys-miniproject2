package utils

import (
	pb "github.com/gustav-mb/disys-miniproject2/chatpb"
	fmt "fmt"
)

type Command struct {
	Name        string
	Description string
}

var Commands = [4]*Command{
	{"help", "Shows all available commands."},
	{"info", "Shows information about the client."},
	{"server", "Shows information about the server"},
	{"exit", "Disconnect from server and exit client."},
}

// Prints a list of available commands
func Help() {
	fmt.Println(Line)
	fmt.Println("HELP")
	fmt.Println("To execute a command, start by writing a forward slash (/) followed by the command name.")
	fmt.Println()
	fmt.Println("AVAILABLE COMMANDS:")
	for _, command := range Commands {
		fmt.Printf("- %v - %v\n", command.Name, command.Description)
	}
	fmt.Println(Line)
}

// Prints the current client info to the console.
func PrintClientInfo(user *pb.User, lamport int32) {
	fmt.Println(Line)
	fmt.Println("Client INFO")
	fmt.Println("ID:", user.Id)
	fmt.Println("Username:", user.Name)
	fmt.Println("Lamport clock:", lamport)
	fmt.Println(Line)
}

func PrintServerInfo(address, port string) {
	fmt.Println(Line)
	fmt.Println("SERVER INFO")
	fmt.Println("IP:", address)
	fmt.Println("Port:", port)
	fmt.Println(Line)
}
