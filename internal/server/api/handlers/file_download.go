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

// FileDownload - checks the validity of the token, save record, upload file on the client
func (h *Handler) FileDownload(ctx context.Context, req *grpc.DownloadBinaryRequest) (*grpc.DownloadBinaryResponse, error) {
	h.logger.Info("file download")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	FileData := &model.FileRequest{}
	FileData.UserID = req.AccessToken.UserId
	FileData.Name = req.Name

	exists, err := h.file.FileExists(FileData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists != true {
		err = errors.ErrFileNotExists
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
