package delivery

import (
	"encoding/json"
	"forum/models"
	"net/http"
	"strconv"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type CreatePostResponse struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Token       string `json:"token"`
	}

	var resp CreatePostResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	post := models.Post{
		Title:       resp.Title,
		Description: resp.Description,
		Category:    []string{resp.Category},
	}
	user, err := h.service.Auth.GetUserByToken(resp.Token)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}
	if err := h.service.Post.CreatePost(post, user.Username); err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))

		return
	}
	h.response(w, statusOK)
}

func (h *Handler) getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodGet {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type GetPostResponse struct {
		Id string `json:"id"`
	}

	var resp GetPostResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	id, err := strconv.Atoi(resp.Id)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}
	post, err := h.service.Post.GetPostById(id)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}
	post.Category, err = h.service.GetCategoriesByPostId(post.Id)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}
	post.Likes, post.Dislikes, err = h.service.Reaction.GetCounts(post.Id, "post")
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}
	if err := h.service.Post.UpdateCountsReactions("post", post.Likes, post.Dislikes, post.Id); err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}
	h.response(w, post)
}
func (h *Handler) getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	posts, err := h.service.GetAllPosts()
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}
	h.response(w, posts)
}
func (h *Handler) getAllComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodGet {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type GetCommentResponse struct {
		Id string `json:"id"`
	}

	var resp GetCommentResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	comments, err := h.service.Comment.GetCommentsByIdPost(resp.Id)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
		return
	}
	h.response(w, comments)
}
