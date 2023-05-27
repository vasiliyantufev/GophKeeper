package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleGetNodeText - get node text
func (h *Handler) HandleGetNodeText(ctx context.Context, req *grpc.GetNodeTextRequest) (*grpc.GetNodeTextResponse, error) {
	TextData := &model.GetNodeTextRequest{}
	TextData.TextId = req.TextId
	GetNodeText, err := h.text.GetNodeText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeTextResponse{}, err
	}
	h.logger.Debug(GetNodeText)
	return &grpc.GetNodeTextResponse{TextId: GetNodeText.ID, Text: GetNodeText.Text}, nil
}
