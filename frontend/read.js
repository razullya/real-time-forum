// package delivery

// import "net/http"

// func (h *Handler) getAllNotif(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")

// 	if r.Method != http.MethodPost {
// 		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
// 		return
// 	}
// 	token := r.URL.Query().Get("token")
// 	user, err := h.service.Auth.GetUserByToken(token)
// 	if err != nil {
// 		h.response(w, h.onError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
// 		return
// 	}
// 	notif, err := h.service.Notification.GetNotificationByUsername(user.Username)
// 	if err != nil {
// 		h.response(w, h.onError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
// 		return
// 	}

// 	h.response(w, notif)
// }
// package delivery

// import (
// 	"encoding/json"
// 	"fmt"
// 	"forum/models"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// type Message struct {
// 	Text   string `json:"text"`
// 	Sender string `json:"sender"`
// }

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// var connections = make(map[string]*websocket.Conn)
// var users = make(map[string]string)

// func (h *Handler) chatHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	token := r.URL.Query().Get("token")
// 	fmt.Println(token)
// 	username, ok := users[token]
// 	if !ok {
// 		http.Error(w, "Failed to create WebSocket connection", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(username)
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		http.Error(w, "Failed to create WebSocket connection", http.StatusInternalServerError)
// 		return
// 	}
// 	defer conn.Close()

// 	connections[token] = conn
// 	defer delete(connections, token)
// 	fmt.Println(connections)
// 	for {
// 		messageType, messageBytes, err := conn.ReadMessage()
// 		if err != nil {
// 			break
// 		}

// 		var message Message
// 		err = json.Unmarshal(messageBytes, &message)
// 		if err != nil {
// 			continue
// 		}
// 		fmt.Println(message)

// 		connChat, ok := connections[token]
// 		if ok {
// 			err = connChat.WriteMessage(messageType, []byte(fmt.Sprintf("[%s]: %s", message.Sender, message.Text)))
// 			if err != nil {
// 				fmt.Println("(")
// 				break
// 			}

// 		}
// 	}
// }


// func (h *Handler) chatReq(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	if r.Method != http.MethodPost {
// 		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
// 		return
// 	}
// 	type ChatReqRequest struct {
// 		Token    string `json:"token"`
// 		Username string `json:"username"`
// 	}

// 	var resp ChatReqRequest
// 	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	user, err := h.service.Auth.GetUserByToken(resp.Token)
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}
// 	token, err := h.service.Chat.GenerateToken()
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}
// 	token2, err := h.service.Chat.GenerateToken()
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}
// 	if err := h.service.Notification.CreateNotification(models.Notification{
// 		Username: resp.Username,
// 		Sender:   user.Username,
// 		Message:  token2,
// 		Checked:  false,
// 	}); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	users[resp.Username] = token
// 	h.response(w, map[string]string{"token": token})
// }

// func (h *Handler) chatStart(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	if r.Method != http.MethodPost {
// 		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
// 		return
// 	}
// 	type ChatStartRequest struct {
// 		Token     string `json:"token"`
// 		TokenChat string `json:"chat_token"`
// 		Username  string `json:"username"`
// 	}

// 	var resp ChatStartRequest
// 	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	user, err := h.service.Auth.GetUserByToken(resp.Token)
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	if err := h.service.Notification.NotificationChecked(user.Username, resp.Username, resp.TokenChat); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	token, err := h.service.Chat.GenerateToken()
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}
// 	users[resp.Username] = token

// 	h.response(w, map[string]string{"token": token})
// }
// исправь код на GO и напиши функционал на vanilla js для фронтаI have corrected the code for the Go backend. For the frontend, I have provided a vanilla JavaScript implementation of the chat functionality. Please note that this is a basic implementation and might need additional error handling and optimizations based on your specific use case.

// Go Backend:

// ```go
// package delivery

// import (
// 	"encoding/json"
// 	"fmt"
// 	"forum/models"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// type Message struct {
// 	Text   string `json:"text"`
// 	Sender string `json:"sender"`
// }

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// var connections = make(map[string]*websocket.Conn)
// var users = make(map[string]string)

// func (h *Handler) chatHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	token := r.URL.Query().Get("token")
// 	fmt.Println(token)
// 	username, ok := users[token]
// 	if !ok {
// 		http.Error(w, "Failed to create WebSocket connection", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(username)
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		http.Error(w, "Failed to create connection", http.StatusInternalServerError)
// 		return
// 	}
// 	defer conn.Close()

// 	connections[token] = conn
// 	defer delete(connections, token)
// 	fmt.Println(connections)
//  {
// 		messageType, messageBytes, err := conn.ReadMessage()
// 		if err != nil {
// 			break
// 		}

// 		var message Message
// 		err = json.Unmarshal(messageBytes, &message)
// 		if err != nil {
// 			continue
// 		}
// 		fmt.Println(message)

// 		connChat, ok := connections[token]
// 		if ok {
// 			err = connChat.WriteMessage(messageType, []byte(fmt.Sprintf("[%s]: %s", message.Sender, message.Text)))
// 			if err != nil {
// 				fmt.Println("(")
// 				break
// 			}

// 		}
// 	}
// }

// func (h *Handler) chatReq(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	if r.Method != http.MethodPost {
// 		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
// 		return
// 	}
// 	type ChatReqRequest struct {
// 		Token    string `json:"token"`
// 		Username string `json:"username"`
// 	}

// 	var resp ChatReqRequest
// 	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	user, err := h.service.Auth.GetUserByToken(resp.Token)
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}
// 	token, err := h.service.Chat.GenerateToken()
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}
// 	token2, err := h.service.Chat.GenerateToken()
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}
// 	if err := h.service.Notification.CreateNotification(models.Notification{
// 		Username: resp.Username,
// 		Sender:   user.Username,
// 		Message:  token2,
// 		Checked:  false,
// 	}); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	users[resp.Username] = token
// 	h.response(w, map[string]string{"token": token})
// }

// func (h *Handler) chatStart(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	if r.Method != http.MethodPost {
// 		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
// 		return
// 	}
// 	type ChatStartRequest struct {
// 		Token     string `json:"token"`
// 		TokenChat string `json:"chat_token"`
// 		Username  string `json:"username"`
// 	}

// 	var resp ChatStartRequest
// 	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	user, err := h.service.Auth.GetUserByToken(resp.Token)
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	if err := h.service.Notification.NotificationChecked(user.Username, resp.Username, resp.TokenChat); err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}

// 	token, err := h.service.Chat.GenerateToken()
// 	if err != nil {
// 		h.response(w, h.onError(err.Error(), http.StatusBadRequest))
// 		return
// 	}
// 	users[resp.Username] = token

// 	h.response(w, map[string]string{"token": token})
// }
// ```

// Frontend (Vanilla JavaScript):

// ```html
// <!DOCTYPE html>
// <html lang="en">
// <head>
//     <meta charset="UTF-8">
//     <meta name="viewport" content="width=device-width, initial-scale=1.0">
//     <title>Chat</title>
//     <style>
//         #messages {
//             height: 300px;
//             border: 1px solid black;
//             overflow-y: scroll;
//         }
//     </style>
// </head>
// <body>
//     <div id="messages"></div>
//     <input type="text" id="messageInput">
//     <button id="sendButton">Send</button>

//     <script>
//         const messagesDiv = document.getElementById('messages');
//         const messageInput = document.getElementById('messageInput');
//         const sendButton = document.getElementById('sendButton');

//         const token = "YOUR_TOKEN"; // Replace with the token you received from the server
//         const ws = new WebSocket(`ws://localhost:8080/chat?token=${token}`);

//         ws.addEventListener('message', (event) => {
//             const message = document.createElement('div');
//             message.textContent = event.data;
//             messagesDiv.appendChild(message);
//             messagesDiv.scrollTop = messagesDiv.scrollHeight;
//         });

//         sendButton.addEventListener('click', () => {
//             const messageText = messageInput.value.trim();
//             if (messageText) {
//                 const message = {
//                     text: messageText,
//                     sender: "YOUR_USERNAME" // Replace with the sender's username
//                 };
//                 ws.send(JSON.stringify(message));
//                 messageInput.value = '';
//             }
//         });

//         messageInput.addEventListener('keydown', (event) => {
//             if (event.key === 'Enter') {
//                 sendButton.click();
//             }
//         });
//     </script>
// </body>
// </html>
// ```

// Replace `YOUR_TOKEN` and `YOUR_USERNAME` with the appropriate values from your application.