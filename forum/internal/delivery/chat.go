package delivery

import (
	"encoding/json"
	"fmt"
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

func (h *Handler) chatHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	fmt.Println(1)
	if err := h.service.Chat.CheckToken([]string{token}); err != nil {
		http.Error(w, "Failed to check token for WebSocket connection", http.StatusInternalServerError)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to create WebSocket connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	connections[token] = conn
	defer delete(connections, token)

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
		// Найдите соединение получателя и отправьте сообщение
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

func (h *Handler) chatCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method != http.MethodPost {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	type ChatCheckRequest struct {
		User  string `json:"username"`
		Token string `json:"token"`
	}
	var req ChatCheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}
	fmt.Println(req)
	token, err := h.service.Chat.GetChatByTokenAndUsername(req.User, req.Token)
	if err != nil {
		if err.Error() == "no user" {
			h.response(w, h.onError(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized))
			return
		}
		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
		return
	}

	h.response(w, map[string]string{"ftoken": token[0], "stoken": token[1]})
}
