package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/iamgoangle/go-graphql-example/cmd/gql/graph/generated"
	"github.com/iamgoangle/go-graphql-example/cmd/gql/graph/model"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo

	usrATodo := model.Todo{
		ID:   "001",
		Text: "test todo",
		Done: false,
		User: &model.User{
			ID:   "usr-001",
			Name: "Teerapong Singthong",
		},
	}

	todos = append(todos, &usrATodo)

	return todos, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
