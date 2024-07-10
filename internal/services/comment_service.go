package services

import (
	"context"
	"ozon-test/internal/db"
	"ozon-test/internal/models"
)

func CreateComment(postID int, userID int, parentCommentID int, content string) (models.Comment, error) {
	query := `INSERT INTO comments (post_id, user_id, parent_comment_id, content) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	var comment models.Comment
	err := db.DB.QueryRow(context.Background(), query, postID, userID, parentCommentID, content).Scan(&comment.ID, &comment.CreatedAt)
	if err != nil {
		return comment, err
	}
	comment.PostID = postID
	comment.UserID = userID
	comment.ParentCommentID = parentCommentID
	comment.Content = content
	return comment, nil
}

func GetCommentsByPostID(postID int) ([]models.Comment, error) {
	query := `SELECT id, post_id, user_id, parent_comment_id, content, created_at FROM comments WHERE post_id=$1`
	rows, err := db.DB.Query(context.Background(), query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.ParentCommentID, &comment.Content, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
