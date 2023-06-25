package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
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

	ListLoginPassword, err := h.loginPassword.GetListLoginPassword(req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list := model.GetListLoginPassword(ListLoginPassword)

	h.logger.Debug(ListLoginPassword)
	return &grpc.GetListLoginPasswordResponse{Node: list}, nil
}
