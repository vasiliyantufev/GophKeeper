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

// HandleDeleteLoginPassword - delete login password
func (h *Handler) HandleDeleteLoginPassword(ctx context.Context, req *grpc.DeleteLoginPasswordRequest) (*grpc.DeleteLoginPasswordResponse, error) {
	h.logger.Info("Delete login password")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteLoginPasswordResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	loginPasswordID, err := h.loginPassword.GetIdLoginPassword(req.Name, req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	metadataRequest := model.DeleteMetadataRequest{loginPasswordID, string(variables.Name), req.Name, string(variables.LoginPassword)}
	err = h.metadata.DeleteMetadata(metadataRequest)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = h.loginPassword.DeleteLoginPassword(loginPasswordID)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DeleteLoginPasswordResponse{Id: loginPasswordID}, nil
}
