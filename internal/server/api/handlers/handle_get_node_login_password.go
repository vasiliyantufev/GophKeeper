package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGetNodeLoginPassword - get node login password
func (h *Handler) HandleGetNodeLoginPassword(ctx context.Context, req *grpc.GetNodeLoginPasswordRequest) (*grpc.GetNodeLoginPasswordResponse, error) {
	h.logger.Info("Get node login password")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetNodeLoginPasswordResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	return &grpc.GetNodeLoginPasswordResponse{}, nil
}
