package delivery

import (
	"encoding/json"
	"fmt"
	"forum/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var statusOK = Status{
	Success: true,
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Status struct {
	Success bool `json:"success"`
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

func (h *Handler) InitRoutes() {
	h.Mux.HandleFunc("/signin", h.signIn)
	h.Mux.HandleFunc("/signup", h.signUp)

	h.Mux.HandleFunc("/logout", h.logOut)

	h.Mux.HandleFunc("/token", h.checkToken)

	h.Mux.HandleFunc("/post", h.getPost)         
	h.Mux.HandleFunc("/post/all", h.getAllPosts) 

	h.Mux.HandleFunc("/post/like", h.likePost)          //
	h.Mux.HandleFunc("/post/dislike", h.dislikePost)    //
	h.Mux.HandleFunc("/post/create", h.createPost)      
	h.Mux.HandleFunc("/post/comments", h.getAllComment) 

	h.Mux.HandleFunc("/comment/create", h.createComment)   
	h.Mux.HandleFunc("/comment/like", h.likeComment)       //
	h.Mux.HandleFunc("/comment/dislike", h.dislikeComment) //

	h.Mux.HandleFunc("/profile", h.getUser) //
	// h.Mux.HandleFunc("/user/other", h.getOtherUser)//

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

		// resp = h.RouterWS(data, path)
		// fmt.Println(resp)
		// fmt.Println("прошли изменения")

		if err := h.ws.WriteMessage(websocket.TextMessage, []byte(resp)); err != nil {
			log.Println(err)
			return
		}
	}
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
