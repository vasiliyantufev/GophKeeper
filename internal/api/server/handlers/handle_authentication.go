package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleAuthentication - authentication user
func (h *Handler) HandleAuthentication(ctx context.Context, req *grpc.AuthenticationRequest) (*grpc.AuthenticationResponse, error) {
	UserData := &model.UserRequest{}
	UserData.Username = req.Username
	UserData.Password = req.Password

	authenticatedUser, err := h.user.Authentication(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.AuthenticationResponse{}, err
	}
	h.logger.Debug(authenticatedUser)
	return &grpc.AuthenticationResponse{UserId: authenticatedUser.ID, Username: authenticatedUser.Username}, nil
}
