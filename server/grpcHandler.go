package main

import (
	"context"
	"log"

	"google.golang.org/grpc/stats"
)

type Handler struct {
}

func (h *Handler) TagRPC(context.Context, *stats.RPCTagInfo) context.Context {
	return context.Background()
}

// HandleRPC processes the RPC stats.
func (h *Handler) HandleRPC(context.Context, stats.RPCStats) {
	
}

func (h *Handler) TagConn(context.Context, *stats.ConnTagInfo) context.Context {
	log.Println("Connected")
	return context.Background()
}

// HandleConn processes the Conn stats.
func (h *Handler) HandleConn(c context.Context, s stats.ConnStats) {
	switch s.(type) {
	case *stats.ConnEnd:
		log.Println("Disconnected")
	}
}
