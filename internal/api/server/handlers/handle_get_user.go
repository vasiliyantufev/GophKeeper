package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleGetUser - get user
func (h *Handler) HandleGetUser(ctx context.Context, req *grpc.GetUserRequest) (*grpc.GetUserResponse, error) {
	var resp string
	return &grpc.GetUserResponse{Resp: resp}, nil
}
