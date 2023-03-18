package delivery

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	vars := mux.Vars(r)
	token := vars["token"]

	user, err := h.service.Auth.GetUserByToken(token)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	fmt.Println(user.Username)
	userResp, err := h.service.User.GetUserByUsername(user.Username)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))

		return
	}
	userResp.Posts, err = h.service.Post.GetPostsByUsername(user.Username)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))

		return
	}

	h.response(w, user)
}

func (h *Handler) getOtherUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	vars := mux.Vars(r)
	username := vars["username"]

	fmt.Println(username)
	userResp, err := h.service.User.GetUserByUsername(username)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))

		return
	}
	userResp.Posts, err = h.service.Post.GetPostsByUsername(username)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))

		return
	}

	h.response(w, userResp)
}
