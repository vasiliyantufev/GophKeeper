package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleCreateUser - create user
func (h *Handler) HandleCreateUser(ctx context.Context, req *grpc.CreateUserRequest) (*grpc.CreateUserResponse, error) {
	var resp string
	return &grpc.CreateUserResponse{Resp: resp}, nil
}
