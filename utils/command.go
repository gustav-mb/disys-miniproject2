package utils

import (
	fmt "fmt"
	pb "chatpb"
)

type Command struct {
	Name        string
	Description string
}

var Commands = [3]*Command{
	{"help", "Shows all available commands."},
	{"info", "Shows the current information about the client."},
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
