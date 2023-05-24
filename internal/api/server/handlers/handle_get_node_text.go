package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleGetMetadata - get metadata
func (h *Handler) HandleGetMetadata(ctx context.Context, req *grpc.GetMetadataRequest) (*grpc.GetMetadataResponse, error) {
	var resp string
	return &grpc.GetMetadataResponse{Resp: resp}, nil
}
