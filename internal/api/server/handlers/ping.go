package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// Ping - checks the database connection
func (h *Handler) Ping(ctx context.Context, req *grpc.PingRequest) (*grpc.PingResponse, error) {
	var resp string
	return &grpc.PingResponse{Resp: resp}, nil
}
