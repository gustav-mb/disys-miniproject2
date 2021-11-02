package utils

import (
	"fmt"
)

const Port = "8080"
const MaxMsgLength = 128
const Line = "--------------------------------------------------------------"

func ShowLogo() {
	fmt.Println(Line)
	fmt.Println("   _____ _     _ _   _                 _____ _           _   ")
	fmt.Println("  / ____| |   (_) | | |               / ____| |         | |  ")
	fmt.Println(" | |    | |__  _| |_| |_ _   _ ______| |    | |__   __ _| |_ ")
	fmt.Println(" | |    | '_ \\| | __| __| | | |______| |    | '_ \\ / _` | __|")
	fmt.Println(" | |____| | | | | |_| |_| |_| |      | |____| | | | (_| | |_ ")
	fmt.Println("  \\_____|_| |_|_|\\__|\\__|\\__, |       \\_____|_| |_|\\__,_|\\__|")
	fmt.Println("                          __/ |                              ")
	fmt.Println("                         |___/                               ")
	fmt.Println(Line)
}
