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

// HandleGetNodeText - get node text
func (h *Handler) HandleGetNodeText(ctx context.Context, req *grpc.GetNodeTextRequest) (*grpc.GetNodeTextResponse, error) {
	h.logger.Info("Get node text")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetNodeTextResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	TextData := &model.GetNodeTextRequest{}
	TextData.UserID = req.AccessToken.UserId
	TextData.Key = string(variables.Name)
	TextData.Value = req.Name
	GetNodeText, err := h.text.GetNodeText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	text := model.GetText(GetNodeText)

	h.logger.Debug(GetNodeText)
	return &grpc.GetNodeTextResponse{Text: text}, nil
}
