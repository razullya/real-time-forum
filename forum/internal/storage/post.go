package storage

import (
	"database/sql"
	"fmt"
	"forum/models"
)

type Post interface {
	//	USER
	CreatePost(creator string, post models.Post) error
	DeletePost(post models.Post) error
	UpdatePost(postId int, post models.Post) error
	UpdateCountsReactions(object string, likes int, dislikes int, postId int) error
	//	GET
	GetPostsByUsername(username string) ([]models.Post, error)
	GetPostById(id int) (models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetAllCategories() ([]string, error)
	GetCategoriesById(id int) ([]string, error)
	//	AUTH
	GetUserByToken(token string) (models.User, error)
}

type PostStorage struct {
	db *sql.DB
}

func newPostStorage(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (p *PostStorage) GetUserByToken(token string) (models.User, error) {
	query := `SELECT id, email, username, hashPassword FROM user WHERE session_token=$1;`
	row := p.db.QueryRow(query, token)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("storage: get user by token: %w", err)
	}
	return user, nil
}

func (p *PostStorage) CreatePost(creator string, post models.Post) error {
	query := `INSERT INTO post(title, description, creator,likes,dislikes) VALUES ($1, $2, $3, $4, $5);`
	res, err := p.db.Exec(query, post.Title, post.Description, creator, 0, 0)
	if err != nil {
		return fmt.Errorf("storage: create post: %w", err)
	}
	postId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("storage: create post: %w", err)
	}
	for i := 0; i < len(post.Category); i++ {
		query = `INSERT INTO categories(tag, id_post) VALUES ($1, $2);`
		if _, err := p.db.Exec(query, post.Category[i], int(postId)); err != nil {
			return fmt.Errorf("storage: create post: %w", err)
		}
	}

	return nil
}

func (p *PostStorage) GetPostsByUsername(username string) ([]models.Post, error) {
	posts := []models.Post{}
	query := `SELECT * FROM post WHERE creator=$1;`

	rows, err := p.db.Query(query, username)
	if err != nil {
		return nil, fmt.Errorf("storage: get all posts by username: %w", err)
	}
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Creator, &post.Title, &post.Description, &post.Likes, &post.Dislikes, &post.CreatedAt); err != nil {
			return nil, fmt.Errorf("storage: get all posts by username: %w", err)
		}
		fmt.Println(len(posts))
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *PostStorage) GetPostById(id int) (models.Post, error) {
	query := `SELECT * FROM post WHERE id=$1;`

	rows := p.db.QueryRow(query, id)

	var post models.Post
	if err := rows.Scan(&post.Id, &post.Creator, &post.Title, &post.Description, &post.Likes, &post.Dislikes, &post.CreatedAt); err != nil {
		return models.Post{}, fmt.Errorf("storage: get post by id: %w", err)
	}

	return post, nil
}

func (p *PostStorage) DeletePost(post models.Post) error {
	query := `DELETE FROM post WHERE id = $1;`
	_, err := p.db.Exec(query, post.Id)
	if err != nil {
		return fmt.Errorf("storage: delete post: %w", err)
	}
	return nil
}

func (p *PostStorage) UpdatePost(postId int, post models.Post) error {
	query := `UPDATE post SET title = $1, description = $2 WHERE id = $3;`
	_, err := p.db.Exec(query, post.Title, post.Description, post.Id)
	if err != nil {
		return fmt.Errorf("storage: delete post: %w", err)
	}
	return nil
}

func (p *PostStorage) GetCategoriesById(id int) ([]string, error) {
	var cats []string
	query := `SELECT tag FROM categories WHERE id_post=$1;`
	row, err := p.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("storage: get categories by id post: %w", err)
	}
	for row.Next() {
		cat := ""
		if err := row.Scan(&cat); err != nil {
			return nil, fmt.Errorf("storage: get categories by id post: %w", err)
		}
		cats = append(cats, cat)
	}
	return cats, nil
}

func (p *PostStorage) GetAllPosts() ([]models.Post, error) {
	posts := []models.Post{}

	query := `SELECT * FROM post;`

	row, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("storage: get all posts: %w", err)
	}
	for row.Next() {
		var post models.Post
		if err := row.Scan(&post.Id, &post.Creator, &post.Title, &post.Description, &post.Likes, &post.Dislikes, &post.CreatedAt); err != nil {
			return nil, fmt.Errorf("storage: get all posts: %w", err)
		}
		post.Category, err = p.GetCategoriesById(post.Id)
		if err != nil {
			return nil, fmt.Errorf("storage: get all posts: %w", err)
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (p *PostStorage) UpdateCountsReactions(object string, likes int, dislikes int, id int) error {
	if object == "post" {
		query := `UPDATE post SET likes =$1,dislikes=$2 WHERE id =$3;`
		_, err := p.db.Exec(query, likes, dislikes, id)
		if err != nil {
			return fmt.Errorf("storage: update counts reactions by post id: %w", err)
		}
	} else if object == "comment" {
		query := `UPDATE comment SET likes =$1,dislikes=$2 WHERE id =$3;`
		_, err := p.db.Exec(query, likes, dislikes, id)
		if err != nil {
			return fmt.Errorf("storage: update counts reactions by comment id: %w", err)
		}
	}

	return nil
}

func (p *PostStorage) GetAllCategories() ([]string, error) {
	var categories []string
	query := `SELECT DISTINCT tag FROM categories;`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("storage: get all categories: %w", err)
	}
	for rows.Next() {
		cat := ""
		if err := rows.Scan(&cat); err != nil {
			return nil, fmt.Errorf("storage: get all categories: %w", err)
		}
		categories = append(categories, cat)
	}
	return categories, nil
}
