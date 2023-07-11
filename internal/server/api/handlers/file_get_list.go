package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FileGetList - file get list
func (h *Handler) FileGetList(ctx context.Context, req *grpc.GetListBinaryRequest) (*grpc.GetListBinaryResponse, error) {
	h.logger.Info("file get list")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetListBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	ListBinary, err := h.binary.GetListBinary(req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list := model.GetListBinary(ListBinary)

	h.logger.Debug(ListBinary)
	return &grpc.GetListBinaryResponse{Node: list}, nil
}
