package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandlePing - checks the database connection
func (h *Handler) HandlePing(ctx context.Context, req *grpc.PingRequest) (*grpc.PingResponse, error) {
	var msg string
	err := h.database.Ping()
	if err != nil {
		msg = "unsuccessful database connection"
		h.logger.Error(err)
		return &grpc.PingResponse{Message: msg}, err
	}
	msg = "successful database connection"
	h.logger.Info(msg)
	return &grpc.PingResponse{Message: msg}, nil
}
