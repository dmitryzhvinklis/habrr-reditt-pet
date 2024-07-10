package resolvers

import (
	"ozon-test/internal/models"
	"ozon-test/internal/services"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"posts": &graphql.Field{
			Type: graphql.NewList(models.PostType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return services.GetPosts()
			},
		},
		"post": &graphql.Field{
			Type: models.PostType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"].(int)
				if ok {
					return services.GetPostByID(id)
				}
				return nil, nil
			},
		},
		"comments": &graphql.Field{
			Type: graphql.NewList(models.CommentType),
			Args: graphql.FieldConfigArgument{
				"post_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				postID, ok := params.Args["post_id"].(int)
				if ok {
					return services.GetCommentsByPostID(postID)
				}
				return nil, nil
			},
		},
	},
})
