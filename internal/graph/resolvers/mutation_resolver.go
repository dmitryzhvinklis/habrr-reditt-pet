package resolvers

import (
	"ozon-test/internal/models"
	"ozon-test/internal/services"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": &graphql.Field{
			Type: models.PostType,
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"allow_comments": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userID := params.Args["user_id"].(int)
				content := params.Args["content"].(string)
				allowComments := params.Args["allow_comments"].(bool)
				return services.CreatePost(userID, content, allowComments)
			},
		},
		"createComment": &graphql.Field{
			Type: models.CommentType,
			Args: graphql.FieldConfigArgument{
				"post_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"parent_comment_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				postID := params.Args["post_id"].(int)
				userID := params.Args["user_id"].(int)
				parentCommentID := params.Args["parent_comment_id"].(int)
				content := params.Args["content"].(string)
				return services.CreateComment(postID, userID, parentCommentID, content)
			},
		},
	},
})
