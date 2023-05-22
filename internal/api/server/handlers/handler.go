package handlers

import (
	grpc "github.com/vasiliyantufev/gophkeeper/internal/api/proto"
)

type Handler struct {
	grpc.UnimplementedDevopsServer
}

// NewHandler - creates a new grpc server instance
func NewHandler() *Handler {
	return &Handler{}
}
