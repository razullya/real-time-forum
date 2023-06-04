package delivery

import (
	"encoding/json"
	"forum/internal/service"
	"log"
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
	log.Println("init routes")

	// router := gin.Default()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.Redirect(http.StatusSeeOther, "/web")
	// })

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"*"},
	// 	AllowHeaders: []string{"*"},
	// }))

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

	h.Mux.HandleFunc("/chat/req", h.chatReq)
	h.Mux.HandleFunc("/chat/start", h.chatStart)

	// запрос  на начала чата (тут он ждет ответ от другого пользователя)
	// уведомление имеется в профиле
	// кнопка начала чата
	h.Mux.HandleFunc("/chat/all", h.getAllDialogs)

	h.Mux.HandleFunc("/chat/ws", h.chatHandler)

	h.Mux.HandleFunc("/otp", h.otp)

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
