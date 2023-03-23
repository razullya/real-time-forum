package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum/internal/service"
	"forum/models"
	"net/http"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	
	type SignInRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var resp SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	sessionToken, err := h.service.Auth.GenerateSessionToken(resp.Username, resp.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			h.response(w, h.onError(err.Error(), http.StatusBadRequest))
			return
		}
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, map[string]string{"token": sessionToken})
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	type SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var resp SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	user := models.User{
		Email:    resp.Email,
		Username: resp.Username,
		Password: resp.Password,
	}
	
	fmt.Println(user)

	if err := h.service.Auth.CreateUser(user); err != nil {
		fmt.Println(err)
		if errors.Is(err, service.ErrInvalidUserName) ||
			errors.Is(err, service.ErrInvalidEmail) ||
			errors.Is(err, service.ErrInvalidPassword) ||
			errors.Is(err, service.ErrUserExist) {
			h.response(w, h.onError(err.Error(), http.StatusBadRequest))
			return
		}
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}

	fmt.Println("no error SignUp")

	h.response(w, statusOK)
}

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	type LogOutRequest struct {
		Token    string `json:"token"`
	}

	var resp LogOutRequest
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	if err := h.service.DeleteSessionToken(resp.Token); err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}

	h.response(w, statusOK)
}

func (h *Handler) checkToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	type CheckTokenRequest struct {
		Token string `json:"token"`
	}

	var resp CheckTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	if err := h.service.CheckToken(resp.Token); err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}

	h.response(w, statusOK)
}
