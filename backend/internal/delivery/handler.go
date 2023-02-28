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
	h.Mux.HandleFunc("/ws", h.wsEndpoint)
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	h.ws = ws
	h.reader(ws)
}
func (h *Handler) reader(conn *websocket.Conn) {
	for {
		_, m, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		var data map[string]interface{}
		if err := json.Unmarshal(m, &data); err != nil {
			// log.Println(err)
			return
		}

		fmt.Println(data)
		resp := h.RouterWS(data)
		fmt.Println(resp)

		if err := h.ws.WriteMessage(websocket.TextMessage, []byte(resp)); err != nil {
			log.Println(err)
			return
		}
	}
}
func (h *Handler) RouterWS(data map[string]interface{}) string {
	action, ok := data["action"].(string)
	if !ok {
		return h.onError("no action")
	}
	fmt.Println(action)
	switch action {
	case "signin":
		return h.signIn(data)
	case "signup":
		return h.signUp(data)
	case "logout":
		return h.logOut(data)
	case "token":
		return h.checkToken(data)
	case "post":
		return h.getPost(data)
	case "post/create":
		return h.createPost(data)
	case "post/all":
		return h.getAllPosts(data)
	case "post/like":
		return h.likePost(data)
	case "post/dislike":
		return h.dislikePost(data)
	case "comment/create":
		return h.createComment(data)
	case "comment/like":
		return h.likeComment(data)
	case "comment/dislike":
		return h.dislikeComment(data)
	case "user":
		return h.getUser(data)
	default:
		return h.onError("bad action")
	}
}
func (h *Handler) structToJSON(data interface{}) ([]byte, error) {
	fmt.Println(data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
