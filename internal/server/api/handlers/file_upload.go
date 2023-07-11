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

// FileUpload - file upload
func (h *Handler) FileUpload(ctx context.Context, req *grpc.UploadBinaryRequest) (*grpc.UploadBinaryResponse, error) {
	h.logger.Info("file upload")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	FileData := &model.FileRequest{}
	FileData.UserID = req.AccessToken.UserId
	FileData.Name = req.Name

	exists, err := h.file.FileExists(FileData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrNameAlreadyExists
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	UploadFile, err := h.file.UploadFile(FileData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = service.UploadFile(h.config.FileFolder, req.AccessToken.UserId, req.Name, req.Data)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(UploadFile.Name)
	return &grpc.UploadBinaryResponse{Name: UploadFile.Name}, nil
}
