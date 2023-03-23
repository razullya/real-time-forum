package delivery

import (
	"fmt"
	"forum/models"
	"net/http"
)

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != http.MethodGet {
		h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
		return
	}
	query := r.URL.Query()
	token := query.Get("token")
	var user models.User
	var err error
	fmt.Println(token, len(token), "token",query.Get("username"))
	if token != "" {

		user, err = h.service.Auth.GetUserByToken(token)
		if err != nil {
			h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))
			return
		}
		user.Password = ""
	} else {
		username := query.Get("username")
		fmt.Println("username")

		user, err = h.service.User.GetUserByUsername(username)
		if err != nil {
			h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))

			return
		}
		user.Posts, err = h.service.Post.GetPostsByUsername(user.Username)
		if err != nil {
			h.response(w, h.onError(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed))

			return
		}
	}

	fmt.Println("111",user)

	h.response(w, user)
}
