package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleLogin - login user
func (h *Handler) HandleLogin(ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	var resp string

	return &grpc.LoginResponse{Resp: resp}, nil
}
