package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleUploadBinary - upload binary data
func (h *Handler) HandleUploadBinary(ctx context.Context, req *grpc.UploadBinaryRequest) (*grpc.UploadBinaryResponse, error) {
	h.logger.Info("Upload binary data")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	err := service.UploadFile(h.config.FileFolder, req.AccessToken.UserId, req.Name, req.Data)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.UploadBinaryResponse{}, nil
}
