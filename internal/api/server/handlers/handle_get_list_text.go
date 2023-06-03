package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGetListText - get list text
func (h *Handler) HandleGetListText(ctx context.Context, req *grpc.GetListTextRequest) (*grpc.GetListTextResponse, error) {
	h.logger.Info("Get list text")

	ListText, err := h.text.GetListText(req.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list := model.GetListData(ListText)

	h.logger.Debug(ListText)
	return &grpc.GetListTextResponse{Node: list}, nil
}
