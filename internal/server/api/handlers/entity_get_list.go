package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EntityGetList - get list entity
func (h *Handler) EntityGetList(ctx context.Context, req *grpc.GetListEntityRequest) (*grpc.GetListEntityResponse, error) {
	h.logger.Info("Get list entity")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	ListEntity, err := h.entity.GetList(req.AccessToken.UserId, req.Type)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list, err := model.GetListEntity(ListEntity)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(list)
	return &grpc.GetListEntityResponse{Node: list}, nil
}
