package delivery

import (
	"encoding/json"
	"fmt"
	"forum/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var statusOK = `{"status":"OK"}`

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Handler struct {
	Mux     *http.ServeMux
	service *service.Service
	ws      *websocket.Conn
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Mux:     http.NewServeMux(),
		service: services,
	}
}

/*
post
-all
-one
-create
-like
-dislike

auth
-signin
-signup
-logout

comment
-create
-like
-dislike

user
-profile
*/
func (h *Handler) InitRoutes() {
	h.Mux.HandleFunc("/signin", h.wsEndpoint)
	h.Mux.HandleFunc("/signup", h.wsEndpoint)

	h.Mux.HandleFunc("/logout", h.wsEndpoint)

	h.Mux.HandleFunc("/token", h.wsEndpoint)

	h.Mux.HandleFunc("/post", h.wsEndpoint)
	h.Mux.HandleFunc("/post/all", h.wsEndpoint)

	h.Mux.HandleFunc("/post/like", h.wsEndpoint)
	h.Mux.HandleFunc("/post/dislike", h.wsEndpoint)
	h.Mux.HandleFunc("/post/create", h.wsEndpoint)

	h.Mux.HandleFunc("/comment/create", h.wsEndpoint)
	h.Mux.HandleFunc("/comment/like", h.wsEndpoint)
	h.Mux.HandleFunc("/comment/dislike", h.wsEndpoint)
	h.Mux.HandleFunc("/user", h.wsEndpoint)
}

func (h *Handler) wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	h.ws = ws
	fmt.Println(r.URL.Path)
	h.reader(ws, r.URL.Path)
}
func (h *Handler) reader(conn *websocket.Conn, path string) {
	for {
		_, m, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		var resp string
		var data map[string]interface{}

		//
		if err := json.Unmarshal(m, &data); err != nil && len(data) != 0 {
			fmt.Println(err)
			return
		}
		// fmt.Println("данные получены")

		resp = h.RouterWS(data, path)
		// fmt.Println(resp)
		// fmt.Println("прошли изменения")

		if err := h.ws.WriteMessage(websocket.TextMessage, []byte(resp)); err != nil {
			log.Println(err)
			return
		}
	}
}
func (h *Handler) RouterWS(data map[string]interface{}, path string) string {

	fmt.Println(path)
	switch path {
	case "/signin":
		return h.signIn(data)
	case "/signup":
		return h.signUp(data)
	case "/logout":
		return h.logOut(data)
	case "/token":
		return h.checkToken(data)
	case "/post":
		return h.getPost(data)
	case "/post/create":
		return h.createPost(data)
	case "/post/all":
		return h.getAllPosts(data)
	case "/post/like":
		return h.likePost(data)
	case "/post/dislike":
		return h.dislikePost(data)
	case "/comment/create":
		return h.createComment(data)
	case "/comment/like":
		return h.likeComment(data)
	case "/comment/dislike":
		return h.dislikeComment(data)
	case "/user":
		return h.getUser(data)
	default:
		return h.onError("bad action")
	}
}
func (h *Handler) structToJSON(data interface{}) ([]byte, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
