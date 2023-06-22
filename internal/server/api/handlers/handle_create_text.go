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

// HandleCreateText - create text
func (h *Handler) HandleCreateText(ctx context.Context, req *grpc.CreateTextRequest) (*grpc.CreateTextResponse, error) {
	h.logger.Info("Create text")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	TextData := &model.CreateTextRequest{}
	TextData.UserID = req.AccessToken.UserId
	TextData.Name = req.Name
	TextData.Description = req.Description
	TextData.Data = req.Text
	if TextData.Name == "" || TextData.Description == "" {
		err := errors.ErrNoMetadataSet
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}
	exists, err := h.text.KeyExists(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrKeyAlreadyExists
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	CreatedText, err := h.text.CreateText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	text := model.GetText(CreatedText)

	Metadata := &model.CreateMetadataRequest{}
	Metadata.EntityId = CreatedText.ID
	Metadata.Key = string(variables.Name)
	Metadata.Value = TextData.Name
	Metadata.Type = string(variables.Text)
	CreatedMetadataName, err := h.metadata.CreateMetadata(Metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	Metadata = &model.CreateMetadataRequest{}
	Metadata.EntityId = CreatedText.ID
	Metadata.Key = string(variables.Description)
	Metadata.Value = TextData.Description
	Metadata.Type = string(variables.Text)
	CreatedMetadataDescription, err := h.metadata.CreateMetadata(Metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(CreatedText)
	h.logger.Debug(CreatedMetadataName)
	h.logger.Debug(CreatedMetadataDescription)
	return &grpc.CreateTextResponse{Text: text}, nil
}
