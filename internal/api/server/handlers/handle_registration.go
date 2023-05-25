package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/service/validator"
)

// HandleRegistration - registration new user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	var resp string

	RegistrationData := &model.RegistrationRequest{}
	RegistrationData.Username = req.Username
	RegistrationData.Password = req.Password

	if correctPassword := validator.VerifyPassword(req.Password); correctPassword == false {
		resp = "password rules: at least 7 letters, 1 number, 1 upper case, 1 special character"
		h.logger.Error(resp)
		return &grpc.RegistrationResponse{Resp: resp}, nil
	}

	_, err := h.user.Registration(RegistrationData)
	if err != nil {
		resp = "unsuccessful registration user"
		h.logger.Error(err)
		return &grpc.RegistrationResponse{Resp: resp}, err
	}
	resp = "successful registration user"
	h.logger.Info(resp)
	return &grpc.RegistrationResponse{Resp: resp}, nil
}
