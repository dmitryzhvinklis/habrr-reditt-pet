package graph

import (
	"ozon-test/internal/graph/resolvers"

	"github.com/graphql-go/graphql"
)

func NewSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        resolvers.QueryType,
		Mutation:     resolvers.MutationType,
		Subscription: resolvers.SubscriptionType, // добавляем подписки
	})
	if err != nil {
		panic(err)
	}
	return schema
}
