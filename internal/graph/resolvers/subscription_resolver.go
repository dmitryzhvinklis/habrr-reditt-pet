package resolvers

import (
	"ozon-test/internal/models"

	"github.com/graphql-go/graphql"
)

var SubscriptionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"newComment": &graphql.Field{
			Type: models.CommentType,
			Args: graphql.FieldConfigArgument{
				"post_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// Реализация подписки на новые комментарии
				// Можно использовать библиотеки, такие как "github.com/graph-gophers/graphql-go" для поддержки подписок
				return nil, nil
			},
		},
	},
})
