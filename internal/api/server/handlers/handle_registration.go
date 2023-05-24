package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleRegistration - registration user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	var resp string

	RegistrationData := &model.RegistrationRequest{}
	RegistrationData.Username = req.Username
	RegistrationData.Password = req.Password

	_, err := h.user.Registration(RegistrationData)
	if err != nil {
		resp = "unsuccessful registration user"
		h.logger.Error(err)
		return &grpc.RegistrationResponse{Resp: resp}, err
	}
	resp = "successful registration user"
	h.logger.Info(resp)

	return &grpc.RegistrationResponse{Resp: resp}, err
}
