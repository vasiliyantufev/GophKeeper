package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/service/validator"
	"github.com/vasiliyantufev/gophkeeper/internal/storage/errors"
)

// HandleRegistration - registration new user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	var resp string

	UserData := &model.UserRequest{}
	UserData.Username = req.Username
	UserData.Password = req.Password

	if correctPassword := validator.VerifyPassword(req.Password); correctPassword != true {
		err := errors.ErrBadPassword
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, err
	}
	exists, err := h.user.UserExists(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, err
	}
	if exists == true {
		err = errors.ErrUserAlreadyExists
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, err
	}

	registeredUser, err := h.user.Registration(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, err
	}
	resp = "successful registration user"
	h.logger.Info(resp)
	h.logger.Debug(registeredUser)
	return &grpc.RegistrationResponse{UserId: registeredUser.ID, Username: registeredUser.Username}, nil
}
