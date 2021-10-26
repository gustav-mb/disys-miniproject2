package chatpb

import fmt "fmt"

const Port = "8080"
const Line = "-------------------------------------------------------"
const MaxMsgLength = 128

type Command struct {
	Name        string
	Description string
}

var Commands = [4]*Command{
	{"help", "Shows all available commands."},
	{"disconnect", "Disconnects you from the server."},
	{"connect", "Reconnects you to the server."},
	{"exit", "Exits the client by permanently disconnecting you from the server."},
}

func ShowLogo() {
	fmt.Println(Line)
	fmt.Println("   _____ _     _ _   _          _____ _           _   ")
	fmt.Println("  / ____| |   (_) | | |        / ____| |         | |  ")
	fmt.Println(" | |    | |__  _| |_| |_ _   _| |    | |__   __ _| |_ ")
	fmt.Println(" | |    | '_ \\| | __| __| | | | |    | '_ \\ / _` | __|")
	fmt.Println(" | |____| | | | | |_| |_| |_| | |____| | | | (_| | |_ ")
	fmt.Println("  \\_____|_| |_|_|\\__|\\__|\\__, |\\_____|_| |_|\\__,_|\\__|")
	fmt.Println("                          __/ |                       ")
	fmt.Println("                         |___/                        ")
	fmt.Println(Line)
}

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
