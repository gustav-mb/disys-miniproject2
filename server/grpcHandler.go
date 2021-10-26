package main

import (
	"context"

	"google.golang.org/grpc/stats"
)

type Handler struct {
}

func (h *Handler) TagRPC(ctx context.Context, stats *stats.RPCTagInfo) context.Context {
	return context.Background()
}

// HandleRPC processes the RPC stats.
func (h *Handler) HandleRPC(ctx context.Context, stats stats.RPCStats) {

}

func (h *Handler) TagConn(ctx context.Context, stats *stats.ConnTagInfo) context.Context {
	return context.Background()
}

// HandleConn processes the Conn stats.
func (h *Handler) HandleConn(ctx context.Context, s stats.ConnStats) {
	switch s.(type) {
	case *stats.ConnEnd:
		//log.Println("Disconnected")
	case *stats.ConnBegin:
		//log.Printf("Connected: %v", ctx)
	}
}
