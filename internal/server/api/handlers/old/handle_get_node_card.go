package old

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGetNodeCard - get node card
func (h *handlers.Handler) HandleGetNodeCard(ctx context.Context, req *grpc.GetNodeCardRequest) (*grpc.GetNodeCardResponse, error) {
	h.logger.Info("Get node card")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetNodeCardResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	CardData := &model.GetNodeCardRequest{}
	CardData.UserID = req.AccessToken.UserId
	CardData.Key = string(variables.Name)
	CardData.Value = req.Name
	GetNodeCard, err := h.card.GetNodeCard(CardData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	card := model.GetCard(GetNodeCard)

	h.logger.Debug(GetNodeCard)
	return &grpc.GetNodeCardResponse{Data: card}, nil
}
