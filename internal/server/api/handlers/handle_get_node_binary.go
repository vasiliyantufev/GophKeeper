package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGetNodeBinary - get node  binary data
func (h *Handler) HandleGetNodeBinary(ctx context.Context, req *grpc.GetNodeBinaryRequest) (*grpc.GetNodeBinaryResponse, error) {
	h.logger.Info("Get node binary data")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetNodeBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	return &grpc.GetNodeBinaryResponse{}, nil
}
