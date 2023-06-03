package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleAuthentication - authentication user
func (h *Handler) HandleAuthentication(ctx context.Context, req *grpc.AuthenticationRequest) (*grpc.AuthenticationResponse, error) {
	h.logger.Info("Authentication")
	UserData := &model.UserRequest{
		Username: req.Username,
		Password: req.Password,
	}

	authenticatedUser, err := h.user.Authentication(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.AuthenticationResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}
	user := model.GetUserData(authenticatedUser)

	h.logger.Debug(authenticatedUser)
	return &grpc.AuthenticationResponse{User: user, AccessToken: "token"}, nil
}
