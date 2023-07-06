package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
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

	BinaryData := &model.BinaryRequest{}
	BinaryData.UserID = req.AccessToken.UserId
	BinaryData.Name = req.Name

	exists, err := h.binary.FileExists(BinaryData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrKeyAlreadyExists
		h.logger.Error(err)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	data, err := service.DownloadFile(h.config.FileFolder, req.AccessToken.UserId, req.Name)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DownloadBinaryResponse{Data: data}, nil
}
