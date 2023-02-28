package delivery

import (
	"forum/models"
)

/*
post
-all
-one
-create
-like
-dislike
*/
func (h *Handler) createPost(data map[string]interface{}) string {
	title, ok := data["title"].(string)
	if !ok {
		return h.onError("no title")
	}
	description, ok := data["description"].(string)
	if !ok {
		return h.onError("no description")
	}
	category, ok := data["category"].(string)
	if !ok {
		return h.onError("no password")
	}
	username, ok := data["username"].(string)
	if !ok {
		return h.onError("no username")
	}
	post := models.Post{
		Title:       title,
		Description: description,
		Category:    []string{category},
	}

	if err := h.service.Post.CreatePost(post, username); err != nil {
		return h.onError(err.Error())
	}
	return statusOK
}
func (h *Handler) getPost(data map[string]interface{}) string {

	id, ok := data["id"].(int)
	if !ok {
		return h.onError("no id")
	}

	post, err := h.service.Post.GetPostById(id)
	if err != nil {
		return h.onError(err.Error())
	}
	comments, err := h.service.GetCommentsByIdPost(id)
	if err != nil {
		return h.onError(err.Error())
	}

	for i := 0; i < len(comments); i++ {
		comments[i].Likes, comments[i].Dislikes, err = h.service.Reaction.GetCounts(comments[i].Id, "comment")
		if err != nil {
			return h.onError(err.Error())
		}
		if err := h.service.Post.UpdateCountsReactions("comment", comments[i].Likes, comments[i].Dislikes, comments[i].Id); err != nil {
			return h.onError(err.Error())
		}
	}
	post.Likes, post.Dislikes, err = h.service.Reaction.GetCounts(post.Id, "post")
	if err != nil {
		return h.onError(err.Error())

	}
	if err := h.service.Post.UpdateCountsReactions("post", post.Likes, post.Dislikes, post.Id); err != nil {
		return h.onError(err.Error())
	}
	resp, err := h.structToJSON(post)
	if err != nil {
		return h.onError(err.Error())
	}
	return string(resp)
}
func (h *Handler) getAllPosts(data map[string]interface{}) string {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		return h.onError(err.Error())

	}
	resp, err := h.structToJSON(posts)
	if err != nil {
		return h.onError(err.Error())
	}
	return string(resp)
}
