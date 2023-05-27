package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/service/validator"
	"github.com/vasiliyantufev/gophkeeper/internal/storage/errors"
)

// HandleCreateText - create text
func (h *Handler) HandleCreateText(ctx context.Context, req *grpc.CreateTextRequest) (*grpc.CreateTextResponse, error) {

	TextData := &model.CreateTextRequest{}
	TextData.UserID = req.UserId
	TextData.Name = req.Name
	TextData.Description = req.Description
	TextData.Text = req.Text

	if TextData.Name == "" {
		err := errors.ErrBadName
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	exists, err := h.text.NameExists(TextData.Name)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	if exists == true {
		err = errors.ErrNameAlreadyExists
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	if correctText := validator.VerifyText(req.Text); correctText != true {
		err := errors.ErrBadText
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}

	Metadata := &model.CreateMetadataRequest{}
	Metadata.Name = TextData.Name
	Metadata.Description = TextData.Description
	CreatedMetadata, err := h.metadata.CreateMetadata(Metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}

	TextData.MetadataID = CreatedMetadata.ID
	CreatedText, err := h.text.CreateText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	h.logger.Debug(CreatedText)

	return &grpc.CreateTextResponse{TextId: CreatedText.ID, Text: CreatedText.Text}, nil
}
