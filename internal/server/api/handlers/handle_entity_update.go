package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleEntityUpdate - update entity
func (h *Handler) HandleEntityUpdate(ctx context.Context, req *grpc.UpdateEntityRequest) (*grpc.UpdateEntityResponse, error) {
	h.logger.Info("Update entity")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.UpdateEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	UpdatedEntityID, err := h.entity.Update(req.AccessToken.UserId, req.Name, req.Type, req.Data)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(UpdatedEntityID)
	return &grpc.UpdateEntityResponse{Id: UpdatedEntityID}, nil
}
