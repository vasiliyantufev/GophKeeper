package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserExist - user exist check
func (h *Handler) UserExist(ctx context.Context, req *grpc.UserExistRequest) (*grpc.UserExistResponse, error) {
	h.logger.Info("user exist check")
	exist, err := h.user.UserExists(req.Username)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UserExistResponse{Exist: false}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	h.logger.Debug(exist)
	return &grpc.UserExistResponse{Exist: exist}, nil
}
