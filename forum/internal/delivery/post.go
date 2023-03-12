package delivery

import (
	"fmt"
	"forum/models"
	"strconv"
)

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
	token, ok := data["token"].(string)
	if !ok {
		return h.onError("no token")
	}
	post := models.Post{
		Title:       title,
		Description: description,
		Category:    []string{category},
	}
	user, err := h.service.Auth.GetUserByToken(token)
	if err != nil {
		return h.onError(err.Error())
	}
	if err := h.service.Post.CreatePost(post, user.Username); err != nil {
		return h.onError(err.Error())
	}
	return statusOK
}
func (h *Handler) getPost(data map[string]interface{}) string {
	fmt.Println(data)
	idIn, ok := data["id"].(string)
	if !ok {
		return h.onError("no id")
	}
	id, err := strconv.Atoi(idIn)
	if err != nil {
		return h.onError("incorrect id")

	}
	post, err := h.service.Post.GetPostById(id)
	if err != nil {
		return h.onError(err.Error())
	}

	post.Category, err = h.service.GetCategoriesByPostId(post.Id)
	if err != nil {
		return h.onError(err.Error())
	}
	// comments, err := h.service.GetCommentsByIdPost(id)
	// if err != nil {
	// 	return h.onError(err.Error())
	// }

	// for i := 0; i < len(comments); i++ {
	// 	comments[i].Likes, comments[i].Dislikes, err = h.service.Reaction.GetCounts(comments[i].Id, "comment")
	// 	if err != nil {
	// 		return h.onError(err.Error())
	// 	}
	// 	if err := h.service.Post.UpdateCountsReactions("comment", comments[i].Likes, comments[i].Dislikes, comments[i].Id); err != nil {
	// 		return h.onError(err.Error())
	// 	}
	// }
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
