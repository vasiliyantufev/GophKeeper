package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
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

	token, err := h.token.Create(user.UserId, 0)
	if err != nil {
		h.logger.Error(err)
		return &grpc.AuthenticationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	h.logger.Debug(authenticatedUser)
	return &grpc.AuthenticationResponse{AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}}, nil
}
