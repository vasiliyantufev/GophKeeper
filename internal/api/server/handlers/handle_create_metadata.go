package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleCreateMetadata - create metadata
func (h *Handler) HandleCreateMetadata(ctx context.Context, req *grpc.CreateMetadataRequest) (*grpc.CreateMetadataResponse, error) {
	var resp string
	return &grpc.CreateMetadataResponse{Resp: resp}, nil
}
