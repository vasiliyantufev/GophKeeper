package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
)

// HandleDeleteLoginPassword - delete text
func (h *Handler) HandleDeleteText(ctx context.Context, req *grpc.DeleteTextRequest) (*grpc.DeleteTextResponse, error) {

	return &grpc.DeleteTextResponse{}, nil
}
