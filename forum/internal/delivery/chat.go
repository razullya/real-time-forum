package delivery

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Text   string `json:"text"`
	Sender string `json:"sender"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = make(map[string]*websocket.Conn)
var users = make(map[string]string)

func (h *Handler) chatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	token := r.URL.Query().Get("token")
	fmt.Println(token)
	username, ok := users[token]
	if !ok {
		http.Error(w, "Failed to create WebSocket connection", http.StatusInternalServerError)
		return
	}
	fmt.Println(username)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to create WebSocket connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	connections[token] = conn
	defer delete(connections, token)
	fmt.Println(connections)
	for {
		messageType, messageBytes, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var message Message
		err = json.Unmarshal(messageBytes, &message)
		if err != nil {
			continue
		}
		fmt.Println(message)

		connChat, ok := connections[token]
		if ok {
			err = connChat.WriteMessage(messageType, []byte(fmt.Sprintf("[%s]: %s", message.Sender, message.Text)))
			if err != nil {
				fmt.Println("(")
				break
			}

		}
	}
}

// func (h *Handler) chatCheck(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	if r.Method != http.MethodPost {
// 		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
// 		return
// 	}
// 	type ChatCheckRequest struct {
// 		User  string `json:"username"`
// 		Token string `json:"token"`
// 	}
// 	var req ChatCheckRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// }

func (h *Handler) chatReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type ChatReqRequest struct {
		Token    string `json:"token"`
		Username string `json:"username"`
	}

	var resp ChatReqRequest
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	user, err := h.service.Auth.GetUserByToken(resp.Token)
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	if err := h.service.Notification.CreateNotification(models.Notification{
		Username: resp.Username,
		Sender:   user.Username,
		Type:     0,
		Message:  "",
		Checked:  false,
	}); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	token, err := h.service.Chat.GenerateToken()
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	users[resp.Username] = token
	h.response(w, map[string]string{"token": token})
}

func (h *Handler) chatStart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type ChatStartRequest struct {
		Token    string `json:"token"`
		Username string `json:"username"`
	}

	var resp ChatStartRequest
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	token, err := h.service.Chat.GenerateToken()
	if err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	users[resp.Username] = token

	h.response(w, map[string]string{"token": token})
}
