package handlers

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleLogin - login user
func (h *Handler) HandleLogin(ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	var resp string

	LoginData := &model.LoginRequest{}
	LoginData.Username = req.Username
	LoginData.Password = req.Password

	exists, err := h.user.Login(LoginData)
	if err != nil {
		resp = "authorisation Error"
		h.logger.Error(err)
		return &grpc.LoginResponse{Resp: resp}, err
	}
	if exists == false {
		resp = "wrong username or password"
		h.logger.Error(err)
		return &grpc.LoginResponse{Resp: resp}, err
	}
	resp = "successful login"
	h.logger.Info(resp)
	return &grpc.LoginResponse{Resp: resp}, nil
}
