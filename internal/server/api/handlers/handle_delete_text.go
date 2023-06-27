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

// HandleDeleteText - delete text
func (h *Handler) HandleDeleteText(ctx context.Context, req *grpc.DeleteTextRequest) (*grpc.DeleteTextResponse, error) {
	h.logger.Info("Delete text")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteTextResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	textID, err := h.text.GetIdText(req.Name, req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	metadataRequest := model.DeleteMetadataRequest{textID, string(variables.Name), req.Name, string(variables.Text)}
	err = h.metadata.DeleteMetadata(metadataRequest)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = h.text.DeleteText(textID)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DeleteTextResponse{Id: textID}, nil
}
