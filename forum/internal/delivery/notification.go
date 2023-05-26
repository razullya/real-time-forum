package delivery

import (
	"fmt"
	"net/http"
)

func (h *Handler) getAllDialogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	token := r.URL.Query().Get("token")
	user, err := h.service.Auth.GetUserByToken(token)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}
	fmt.Println(user.Username)
	notif, err := h.service.Notification.GetNotificationByUsername(user.Username)
	if err != nil {
		fmt.Println(err)
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}
	fmt.Println(notif)

	h.response(w, notif)
}
