package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandlePing - checks the database connection
func (h *Handler) HandlePing(ctx context.Context, req *grpc.PingRequest) (*grpc.PingResponse, error) {
	var resp string
	err := h.database.Ping()
	if err != nil {
		resp = "unsuccessful database connection"
		h.logger.Error(err)
		return &grpc.PingResponse{Resp: resp}, err
	}
	resp = "successful database connection"
	h.logger.Info(resp)
	return &grpc.PingResponse{Resp: resp}, nil
}
