package delivery

func (h *Handler) getUser(data map[string]interface{}) string {
	username, ok := data["username"].(string)
	if !ok {
		return h.onError("no username")
	}
	user, err := h.service.User.GetUserByUsername(username)
	if err != nil {
		return h.onError(err.Error())
	}
	user.Posts, err = h.service.Post.GetPostsByUsername(username)
	if err != nil {
		return h.onError(err.Error())
	}
	resp, err := h.structToJSON(user)
	if err != nil {
		return h.onError(err.Error())
	}
	return string(resp)
}
