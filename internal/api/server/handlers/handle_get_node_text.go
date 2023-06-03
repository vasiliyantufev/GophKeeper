package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGetNodeText - get node text
func (h *Handler) HandleGetNodeText(ctx context.Context, req *grpc.GetNodeTextRequest) (*grpc.GetNodeTextResponse, error) {
	h.logger.Info("Get node text")

	TextData := &model.GetNodeTextRequest{}
	TextData.UserID = req.UserId
	TextData.Key = req.Key
	TextData.Value = req.Value
	GetNodeText, err := h.text.GetNodeText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	text := model.GetTextData(GetNodeText)

	h.logger.Debug(GetNodeText)
	return &grpc.GetNodeTextResponse{Text: text}, nil
}
