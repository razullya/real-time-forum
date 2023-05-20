package delivery

import "net/http"

func (h *Handler) getAllNotif(w http.ResponseWriter, r *http.Request) {
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
	notif, err := h.service.Notification.GetNotificationByUsername(user.Username)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}

	h.response(w, notif)
}
