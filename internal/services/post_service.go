package services

import (
	"context"
	"ozon-test/internal/db"
	"ozon-test/internal/models"
)

func CreatePost(userID int, content string, allowComments bool) (models.Post, error) {
	query := `INSERT INTO posts (user_id, content, allow_comments) VALUES ($1, $2, $3) RETURNING id`
	var post models.Post
	err := db.DB.QueryRow(context.Background(), query, userID, content, allowComments).Scan(&post.ID)
	if err != nil {
		return post, err
	}
	post.UserID = userID
	post.Content = content
	post.AllowComments = allowComments
	return post, nil
}

func GetPosts() ([]models.Post, error) {
	query := `SELECT id, user_id, content, allow_comments FROM posts`
	rows, err := db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.AllowComments)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostByID(id int) (models.Post, error) {
	query := `SELECT id, user_id, content, allow_comments FROM posts WHERE id=$1`
	var post models.Post
	err := db.DB.QueryRow(context.Background(), query, id).Scan(&post.ID, &post.UserID, &post.Content, &post.AllowComments)
	if err != nil {
		return post, err
	}
	return post, nil
}
