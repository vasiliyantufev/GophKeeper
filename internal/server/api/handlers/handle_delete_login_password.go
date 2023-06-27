package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

// HandleDeleteLoginPassword - delete login password
func (h *Handler) HandleDeleteLoginPassword(ctx context.Context, req *grpc.CreateLoginPasswordRequest) (*grpc.DeleteLoginPasswordResponse, error) {

	return &grpc.DeleteLoginPasswordResponse{}, nil
}
