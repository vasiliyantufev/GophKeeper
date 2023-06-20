package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleCreateCard - create card
func (h *Handler) HandleCreateCard(ctx context.Context, req *grpc.CreateCardRequest) (*grpc.CreateCardResponse, error) {
	h.logger.Info("Create card")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.CreateCardResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	CardData := &model.CreateCardRequest{}
	CardData.UserID = req.AccessToken.UserId
	CardData.Name = req.Name
	CardData.CardData = req.Data
	if CardData.Name == "" {
		err := errors.ErrNoMetadataSet
		h.logger.Error(err)
		return &grpc.CreateCardResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}
	exists, err := h.card.KeyExists(CardData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrKeyAlreadyExists
		h.logger.Error(err)
		return &grpc.CreateCardResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	CreatedCard, err := h.card.CreateCard(CardData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	card := model.GetCardData(CreatedCard)

	Metadata := &model.CreateMetadataRequest{}
	Metadata.EntityId = CreatedCard.ID
	Metadata.Key = string(variables.Name)
	Metadata.Value = CardData.Name
	Metadata.Type = string(variables.Card)
	CreatedMetadataName, err := h.metadata.CreateMetadata(Metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(CreatedCard)
	h.logger.Debug(CreatedMetadataName)
	return &grpc.CreateCardResponse{Data: card}, nil
}
