package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleCreateCart - create card
func (h *Handler) HandleCreateCard(ctx context.Context, req *grpc.CreateCardRequest) (*grpc.CreateCardResponse, error) {
	h.logger.Info("Create cart")

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
	CardData.Card = req.Data
	if CardData.Name == "" {
		err := errors.ErrNoMetadataSet
		h.logger.Error(err)
		return &grpc.CreateCardResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}
	//exists, err := h.text.KeyExists(TextData)
	//if err != nil {
	//	h.logger.Error(err)
	//	return &grpc.CreateTextResponse{}, status.Errorf(
	//		codes.Internal, err.Error(),
	//	)
	//}
	//if exists == true {
	//	err = errors.ErrKeyAlreadyExists
	//	h.logger.Error(err)
	//	return &grpc.CreateTextResponse{}, status.Errorf(
	//		codes.AlreadyExists, err.Error(),
	//	)
	//}

	//CreatedCard, err := h.text.CreateText(CardData)
	//if err != nil {
	//	h.logger.Error(err)
	//	return &grpc.CreateCardResponse{}, status.Errorf(
	//		codes.Internal, err.Error(),
	//	)
	//}
	//text := model.GetTextData(CreatedText)
	//
	//Metadata := &model.CreateMetadataRequest{}
	//Metadata.EntityId = CreatedText.ID
	//Metadata.Key = string(variables.Name)
	//Metadata.Value = TextData.Name
	//Metadata.Type = string(variables.Text)
	//CreatedMetadataName, err := h.metadata.CreateMetadata(Metadata)
	//if err != nil {
	//	h.logger.Error(err)
	//	return &grpc.CreateCardResponse{}, status.Errorf(
	//		codes.Internal, err.Error(),
	//	)
	//}
	//
	//h.logger.Debug(CreatedText)
	//h.logger.Debug(CreatedMetadataName)
	return &grpc.CreateCardResponse{}, nil
}
