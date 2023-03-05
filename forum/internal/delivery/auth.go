package delivery

import (
	"errors"
	"fmt"
	"forum/internal/service"
	"forum/models"
)

func (h *Handler) signIn(data map[string]interface{}) string {
	username, ok := data["username"].(string)
	if !ok {
		return h.onError("no one tokens")
	}
	password, ok := data["password"].(string)
	if !ok {
		return h.onError("no one tokens")
	}

	sessionToken, err := h.service.Auth.GenerateSessionToken(username, password)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {

			return h.onError(err.Error())
		}
		return h.onError(err.Error())
	}

	return fmt.Sprintf(`{"token": "%s"}`, sessionToken)
}

func (h *Handler) signUp(data map[string]interface{}) string {
	email, ok := data["email"].(string)
	if !ok {
		return h.onError("no email")
	}
	username, ok := data["username"].(string)
	if !ok {
		return h.onError("no username")
	}
	password, ok := data["password"].(string)
	if !ok {
		return h.onError("no password")
	}
	user := models.User{
		Email:    email,
		Username: username,
		Password: password,
	}
	fmt.Println(user)
	if err := h.service.Auth.CreateUser(user); err != nil {
		if errors.Is(err, service.ErrInvalidUserName) ||
			errors.Is(err, service.ErrInvalidEmail) ||
			errors.Is(err, service.ErrInvalidPassword) ||
			errors.Is(err, service.ErrUserExist) {
			return h.onError(err.Error())

		}

		return h.onError(err.Error())
	}
	fmt.Println(user)
	return statusOK
}

func (h *Handler) logOut(data map[string]interface{}) string {
	token, ok := data["token"].(string)
	if !ok {
		return h.onError("no one tokens")
	}
	if err := h.service.DeleteSessionToken(token); err != nil {
		return h.onError(err.Error())
	}
	return statusOK
}

func (h *Handler) checkToken(data map[string]interface{}) string {
	token, ok := data["token"].(string)
	if !ok {
		return h.onError("no one tokens")
	}
	fmt.Println(token)
	if err := h.service.CheckToken(token); err != nil {
		return h.onError("incorrect token")
	}
	return statusOK
}
