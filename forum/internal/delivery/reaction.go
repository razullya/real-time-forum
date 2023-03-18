package delivery

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	vars := mux.Vars(r)
	id := vars["post_id"]
	token := vars["token"]
	postId, err := strconv.Atoi(id)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}
	user, err := h.service.Auth.GetUserByToken(token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Reaction.CreateReaction(postId, "post", "like", user.Username); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}
	h.response(w, statusOK)

}
func (h *Handler) dislikePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	vars := mux.Vars(r)
	id := vars["post_id"]
	token := vars["token"]
	postId, err := strconv.Atoi(id)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}
	user, err := h.service.Auth.GetUserByToken(token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Reaction.CreateReaction(postId, "post", "dislike", user.Username); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}
	h.response(w, statusOK)
}
func (h *Handler) likeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	vars := mux.Vars(r)
	id := vars["post_id"]
	token := vars["token"]
	postId, err := strconv.Atoi(id)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}
	user, err := h.service.Auth.GetUserByToken(token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Reaction.CreateReaction(postId, "comment", "like", user.Username); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}
	h.response(w, statusOK)
}
func (h *Handler) dislikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	vars := mux.Vars(r)
	id := vars["post_id"]
	token := vars["token"]
	postId, err := strconv.Atoi(id)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}
	user, err := h.service.Auth.GetUserByToken(token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Reaction.CreateReaction(postId, "comment", "dislike", user.Username); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}
	h.response(w, statusOK)
}

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	token := vars["token"]
	description := vars["description"]

	if err := h.service.Comment.CreateComment(id, token, description); err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	h.response(w, statusOK)
}
