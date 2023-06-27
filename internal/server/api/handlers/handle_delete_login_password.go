package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleDeleteLoginPassword - delete login password
func (h *Handler) HandleDeleteLoginPassword(ctx context.Context, req *grpc.DeleteLoginPasswordRequest) (*grpc.DeleteLoginPasswordResponse, error) {
	h.logger.Info("Delete login password")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteLoginPasswordResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}
	return &grpc.DeleteLoginPasswordResponse{}, nil
}
