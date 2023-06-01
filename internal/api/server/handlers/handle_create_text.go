package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/storage/errors"
)

// HandleCreateText - create text
func (h *Handler) HandleCreateText(ctx context.Context, req *grpc.CreateTextRequest) (*grpc.CreateTextResponse, error) {
	h.logger.Info("Create text")

	TextData := &model.CreateTextRequest{}
	TextData.UserID = req.UserId
	TextData.Key = req.Key
	TextData.Value = req.Value
	TextData.Text = req.Text

	if TextData.Key == "" || TextData.Value == "" {
		err := errors.ErrNoMetadataSet
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	exists, err := h.text.KeyExists(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	if exists == true {
		err = errors.ErrKeyAlreadyExists
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}

	CreatedText, err := h.text.CreateText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	text := model.GetTextData(CreatedText)

	Metadata := &model.CreateMetadataRequest{}
	Metadata.EntityId = CreatedText.ID
	Metadata.Key = TextData.Key
	Metadata.Value = TextData.Value
	Metadata.Type = TextData.Type
	CreatedMetadata, err := h.metadata.CreateMetadata(Metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}

	h.logger.Debug(CreatedText)
	h.logger.Debug(CreatedMetadata)
	return &grpc.CreateTextResponse{Text: text}, nil
}
