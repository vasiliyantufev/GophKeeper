package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleGetToken - get token
func (h *Handler) HandleGetToken(ctx context.Context, req *grpc.GetTokenRequest) (*grpc.GetTokenResponse, error) {
	var token string
	return &grpc.GetTokenResponse{Token: token}, nil
}
