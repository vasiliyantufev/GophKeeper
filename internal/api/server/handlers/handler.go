package handlers

import (
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

type Handler struct {
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler() *Handler {
	return &Handler{}
}
