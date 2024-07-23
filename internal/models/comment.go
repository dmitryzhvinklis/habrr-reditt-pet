package models

import (
	"time"

	"github.com/graphql-go/graphql"
)

var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.Int,
		},
		"PostID": &graphql.Field{
			Type: graphql.Int,
		},
		"UserID": &graphql.Field{
			Type: graphql.Int,
		},
		"ParentCommentID": &graphql.Field{
			Type: graphql.Int,
		},
		"Content": &graphql.Field{
			Type: graphql.String,
		},
		"CreatedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

type Comment struct {
	ID              int
	PostID          int
	UserID          int
	ParentCommentID int
	Content         string
	CreatedAt       time.Time
}
