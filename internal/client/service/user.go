package service

import "github.com/vasiliyantufev/gophkeeper/internal/client/model"

var userId int64
var accessToken string

func UserExist(username string) bool {
	return false
}

func Authentication(username, password string) (model.User, bool) {
	userId = 1
	accessToken = "token"
	user := model.User{ID: userId, Name: username, Password: password, AccessToken: accessToken}
	return user, true
}

func Registration(username, password string) model.User {
	userId = 1
	accessToken = "token"
	user := model.User{ID: userId, Name: username, Password: password, AccessToken: accessToken}
	return user
}
