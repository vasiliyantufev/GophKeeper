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

// FileRemove - file remove
func (h *Handler) FileRemove(ctx context.Context, req *grpc.DeleteBinaryRequest) (*grpc.DeleteBinaryResponse, error) {
	h.logger.Info("file remove")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	BinaryData := &model.BinaryRequest{}
	BinaryData.UserID = req.AccessToken.UserId
	BinaryData.Name = req.Name

	BinaryId, err := h.binary.DeleteBinary(BinaryData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = service.RemoveFile(h.config.FileFolder, req.AccessToken.UserId, req.Name)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DeleteBinaryResponse{Id: BinaryId}, nil
}
