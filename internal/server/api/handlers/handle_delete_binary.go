package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleDeleteBinary - delete  binary data
func (h *Handler) HandleDeleteBinary(ctx context.Context, req *grpc.DeleteBinaryRequest) (*grpc.DeleteBinaryResponse, error) {
	h.logger.Info("Delete binary data")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	return &grpc.DeleteBinaryResponse{}, nil
}
