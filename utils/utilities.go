package utils

import (
	pb "chatpb"
	fmt "fmt"
	"time"
)

const Port = "8080"
const MaxMsgLength = 128
const Line = "-------------------------------------------------------"

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

func NewMessage(user *pb.User, content string, lamport int32) *pb.Message {
	return &pb.Message{
		User: user,
		Content: content,
		Timestamp: time.Now().String(),
		Lamport: lamport,
	}
}

func MaxLamport(x, y int32) int32 {
	if x < y {
		return y
	}

	return x
}
