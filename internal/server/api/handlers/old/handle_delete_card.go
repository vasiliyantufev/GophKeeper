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

// HandleDeleteCard - delete card
func (h *handlers.Handler) HandleDeleteCard(ctx context.Context, req *grpc.DeleteCardRequest) (*grpc.DeleteCardResponse, error) {
	h.logger.Info("Delete card")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteCardResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	cardID, err := h.card.GetIdCard(req.Name, req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	metadataRequest := model.DeleteMetadataRequest{cardID, string(variables.Name), req.Name, string(variables.Card)}
	err = h.metadata.DeleteMetadata(metadataRequest)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = h.card.DeleteCard(cardID)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DeleteCardResponse{Id: cardID}, nil
}
