package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGetListLoginPassword - get list login password
func (h *Handler) HandleGetListLoginPassword(ctx context.Context, req *grpc.GetListLoginPasswordRequest) (*grpc.GetListLoginPasswordResponse, error) {
	h.logger.Info("Get list login password")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetListLoginPasswordResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}
	return &grpc.GetListLoginPasswordResponse{}, nil
}
