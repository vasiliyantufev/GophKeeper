package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleGetData - get data
func (h *Handler) HandleGetData(ctx context.Context, req *grpc.GetDataRequest) (*grpc.GetDataResponse, error) {
	var resp string
	return &grpc.GetDataResponse{Resp: resp}, nil
}
