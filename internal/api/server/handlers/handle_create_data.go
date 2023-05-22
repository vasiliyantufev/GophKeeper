package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleCreateData - create data
func (h *Handler) HandleCreateData(ctx context.Context, req *grpc.CreateDataRequest) (*grpc.CreateDataResponse, error) {
	var resp string
	return &grpc.CreateDataResponse{Resp: resp}, nil
}
