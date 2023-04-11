package delivery

import (
	"encoding/json"
	"errors"
	"forum/internal/service"
	"forum/models"
	"net/http"
	"net/smtp"
	"strings"
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

	_, err := h.service.Auth.GenerateSessionToken(resp.Username, resp.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			h.response(w, h.onError(err.Error(), http.StatusBadRequest))
			return
		}
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}
	from := "eurasian@internet.ru"
	password := "vJvdasdR6jmkPYcEsNh5"
	to := "luap11@mail.ru"
	subject := "OTP"
	body := "КОД СГЕНЕРИРОВАННЫЙ"

	err = sendMail(from, password, to, subject, body)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}

	h.response(w, statusOK)
}
func sendMail(from, password, to, subject, body string) error {
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, strings.Split(to, ","), message)
	return err
}
func (h *Handler) otp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type OTPRequest struct {
		Code string `json:"code"`
	}

	var resp OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

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

	if err := h.service.Auth.CreateUser(user); err != nil {
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
		Token string `json:"token"`
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

	if err := h.service.Auth.CheckToken(resp.Token); err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}

	h.response(w, statusOK)
}
