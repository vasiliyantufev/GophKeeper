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

	if correctText := validator.VerifyText(req.Text); correctText != true {
		err := errors.ErrBadText
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}

	TextData := &model.CreateTextRequest{}
	TextData.UserID = req.UserId
	TextData.Text = req.Text

	CreatedText, err := h.text.CreateText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	h.logger.Debug(CreatedText)

	return &grpc.CreateTextResponse{TextId: CreatedText.ID, Text: CreatedText.Text}, nil
}
