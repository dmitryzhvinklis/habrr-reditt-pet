package models

import (
	"github.com/graphql-go/graphql"
)

var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id":                &graphql.Field{Type: graphql.Int},
		"post_id":           &graphql.Field{Type: graphql.Int},
		"user_id":           &graphql.Field{Type: graphql.Int},
		"parent_comment_id": &graphql.Field{Type: graphql.Int},
		"content":           &graphql.Field{Type: graphql.String},
		"created_at":        &graphql.Field{Type: graphql.String},
	},
})

type Comment struct {
	ID              int
	PostID          int
	UserID          int
	ParentCommentID int
	Content         string
	CreatedAt       string
}
