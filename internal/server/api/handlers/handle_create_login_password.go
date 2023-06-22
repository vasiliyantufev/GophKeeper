package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleCreateLoginPassword - create login password
func (h *Handler) HandleCreateLoginPassword(ctx context.Context, req *grpc.CreateLoginPasswordRequest) (*grpc.CreateLoginPasswordResponse, error) {
	h.logger.Info("Create login password")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.CreateLoginPasswordResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	return &grpc.CreateLoginPasswordResponse{}, nil
}
