package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleLogin - login user
func (h *Handler) HandleLogin(ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	var resp string

	UserData := &model.UserRequest{}
	UserData.Username = req.Username
	UserData.Password = req.Password

	authentication, err := h.user.Authentication(UserData)
	if err != nil {
		resp = "server error"
		h.logger.Error(err)
		return &grpc.LoginResponse{Resp: resp}, err
	}
	if authentication == false {
		resp = "wrong username or password"
		h.logger.Error(err)
		return &grpc.LoginResponse{Resp: resp}, nil
	}

	resp = "successful login"
	h.logger.Info(resp)
	return &grpc.LoginResponse{Resp: resp}, nil
}
