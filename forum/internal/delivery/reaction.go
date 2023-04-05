package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type LikePostRequest struct {
		Token string `json:"token"`
		Id    string `json:"id"`
	}

	var req LikePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	fmt.Println(req)
	user, err := h.service.Auth.GetUserByToken(req.Token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	id, err := strconv.Atoi(req.Id)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	if err := h.service.Reaction.CreateReaction(id, "post", "like", user.Username); err != nil {
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
	type DislikePostRequest struct {
		Token string `json:"token"`
		Id    string `json:"id"`
	}

	var req DislikePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	user, err := h.service.Auth.GetUserByToken(req.Token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Reaction.CreateReaction(id, "post", "dislike", user.Username); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}
	h.response(w, statusOK)
}
func (h *Handler) likeComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type LikeCommentRequest struct {
		Token string `json:"token"`
		Id    string `json:"id"`
	}

	var req LikeCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	user, err := h.service.Auth.GetUserByToken(req.Token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Reaction.CreateReaction(id, "comment", "like", user.Username); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusInternalServerError))
		return
	}
	h.response(w, statusOK)
}
func (h *Handler) dislikeComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type DislikeCommentRequest struct {
		Token string `json:"token"`
		Id    string `json:"id"`
	}

	var req DislikeCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	user, err := h.service.Auth.GetUserByToken(req.Token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	if err := h.service.Reaction.CreateReaction(id, "comment", "dislike", user.Username); err != nil {
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
	type CreateCommentRequest struct {
		Id          string `json:"id"`
		Token       string `json:"token"`
		Description string `json:"description"`
	}

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	fmt.Println(req)
	if err := h.service.Comment.CreateComment(req.Id, req.Token, req.Description); err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	h.response(w, statusOK)
}
