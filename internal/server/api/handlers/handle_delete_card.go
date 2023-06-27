package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

// HandleDeleteCard - delete card
func (h *Handler) HandleDeleteCard(ctx context.Context, req *grpc.DeleteCardRequest) (*grpc.DeleteCardResponse, error) {

	return &grpc.DeleteCardResponse{}, nil
}
