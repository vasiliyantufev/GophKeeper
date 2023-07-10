package handlers

import (
	"context"
	"encoding/json"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleEntityCreate - create entity
func (h *Handler) HandleEntityCreate(ctx context.Context, req *grpc.CreateEntityRequest) (*grpc.CreateEntityResponse, error) {
	h.logger.Info("Create entity")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	var metadata model.MetadataEntity
	err := json.Unmarshal([]byte(req.Metadata), &metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, err
	}

	EntityData := &model.CreateEntityRequest{}
	EntityData.UserID = req.AccessToken.UserId
	EntityData.Data = req.Data
	EntityData.Metadata = metadata
	if metadata.Name == "" {
		err := errors.ErrNoMetadataSet
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}

	exists, err := h.entity.Exists(EntityData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrNameAlreadyExists
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	CreatedEntity, err := h.entity.Create(EntityData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.CreateEntityResponse{Id: CreatedEntity.ID}, nil
}
