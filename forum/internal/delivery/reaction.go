package delivery

func (h *Handler) likePost(data map[string]interface{}) string {
	postId, ok := data["post_id"].(int)
	if !ok {
		return h.onError("no id")
	}
	username, ok := data["username"].(string)
	if !ok {
		return h.onError("no username")
	}
	if err := h.service.Reaction.CreateReaction(postId, "post", "like", username); err != nil {
		return h.onError(err.Error())
	}

	return statusOK
}
func (h *Handler) dislikePost(data map[string]interface{}) string {
	postId, ok := data["post_id"].(int)
	if !ok {
		return h.onError("no id")
	}
	username, ok := data["username"].(string)
	if !ok {
		return h.onError("no username")
	}
	if err := h.service.Reaction.CreateReaction(postId, "post", "dislike", username); err != nil {
		return h.onError(err.Error())
	}
	return statusOK
}
func (h *Handler) likeComment(data map[string]interface{}) string {
	commentId, ok := data["comment_id"].(int)
	if !ok {
		return h.onError("no id")
	}
	username, ok := data["username"].(string)
	if !ok {
		return h.onError("no username")
	}
	if err := h.service.Reaction.CreateReaction(commentId, "comment", "like", username); err != nil {
		return h.onError(err.Error())

	}
	return statusOK
}
func (h *Handler) dislikeComment(data map[string]interface{}) string {
	commentId, ok := data["comment_id"].(int)
	if !ok {
		return h.onError("no id")
	}
	username, ok := data["username"].(string)
	if !ok {
		return h.onError("no username")
	}
	if err := h.service.Reaction.CreateReaction(commentId, "comment", "dislike", username); err != nil {
		return h.onError(err.Error())

	}
	return statusOK
}

func (h *Handler) createComment(data map[string]interface{}) string {
	id, ok := data["id"].(int)
	if !ok {
		return h.onError("no id")
	}
	username, ok := data["user"].(string)
	if !ok {
		return h.onError("no username")
	}
	comment, ok := data["comment"].(string)
	if !ok {
		return h.onError("no comment")
	}

	if err := h.service.Comment.CreateComment(id, username, comment); err != nil {
		return h.onError(err.Error())
	}
	return statusOK
}
