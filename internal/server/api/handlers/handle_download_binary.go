package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleDownloadBinary - Download  binary data
func (h *Handler) HandleDownloadBinary(ctx context.Context, req *grpc.DownloadBinaryRequest) (*grpc.DownloadBinaryResponse, error) {
	h.logger.Info("Download binary data")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	return &grpc.DownloadBinaryResponse{}, nil
}
