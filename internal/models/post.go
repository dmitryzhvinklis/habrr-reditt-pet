package models

import (
	"github.com/graphql-go/graphql"
)

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.Int},
		"user_id":        &graphql.Field{Type: graphql.Int},
		"content":        &graphql.Field{Type: graphql.String},
		"allow_comments": &graphql.Field{Type: graphql.Boolean},
	},
})

type Post struct {
	ID            int
	UserID        int
	Content       string
	AllowComments bool
}
