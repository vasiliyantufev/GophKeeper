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

// HandleRegistration - registration new user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	h.logger.Info("Registration")

	UserData := &model.UserRequest{}
	UserData.Username = req.Username
	UserData.Password = req.Password

	exists, err := h.user.UserExists(UserData.Username)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrUsernameAlreadyExists
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}
	registeredUser, err := h.user.Registration(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	user := model.GetUserData(registeredUser)

	token, err := h.token.Create(user.UserId, h.config.AccessTokenLifetime)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)

	err = service.CreateStorageUser(h.config.FileFolder, token.UserID)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(registeredUser)
	return &grpc.RegistrationResponse{AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
		CreatedAt: created, EndDateAt: endDate}}, nil
}
