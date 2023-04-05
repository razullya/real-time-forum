package delivery

import (
	"encoding/json"
	"fmt"
	"forum/internal/service"
	"net/http"
)

var statusOK = Status{
	Success: true,
}

type Status struct {
	Success bool `json:"success"`
}

type Handler struct {
	Mux     *http.ServeMux
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Mux:     http.NewServeMux(),
		service: services,
	}
}

func (h *Handler) InitRoutes() {
	h.Mux.HandleFunc("/signin", h.signIn)
	h.Mux.HandleFunc("/signup", h.signUp)

	h.Mux.HandleFunc("/logout", h.logOut)

	h.Mux.HandleFunc("/token", h.checkToken)

	h.Mux.HandleFunc("/post", h.getPost)
	h.Mux.HandleFunc("/post/all", h.getAllPosts)

	h.Mux.HandleFunc("/post/like", h.likePost)       //
	h.Mux.HandleFunc("/post/dislike", h.dislikePost) //
	h.Mux.HandleFunc("/post/create", h.createPost)
	h.Mux.HandleFunc("/post/comments", h.getAllComment)

	h.Mux.HandleFunc("/comment/create", h.createComment)
	h.Mux.HandleFunc("/comment/like", h.likeComment)       //
	h.Mux.HandleFunc("/comment/dislike", h.dislikeComment) //

	h.Mux.HandleFunc("/profile", h.getUser) //
	h.Mux.HandleFunc("/chat/check", h.chatCheck)
	h.Mux.HandleFunc("/chat/create", h.chatCreate)


	h.Mux.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("2222")
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		serveWs(w, r)
	})

}


func (h *Handler) response(w http.ResponseWriter, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
