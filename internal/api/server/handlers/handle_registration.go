package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleRegistration - registration user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {

	RegistrationData := &model.RegistrationRequest{}

	RegistrationData.Username = req.Username
	RegistrationData.Password = req.Password

	var resp string
	return &grpc.RegistrationResponse{Resp: resp}, nil
}
