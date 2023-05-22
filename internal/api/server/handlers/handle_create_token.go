package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleCreateToken - create token
func (h *Handler) HandleCreateToken(ctx context.Context, req *grpc.CreateTokenRequest) (*grpc.CreateTokenResponse, error) {
	var token string
	return &grpc.CreateTokenResponse{Token: token}, nil
}
