package graph

import "github.com/iamgoangle/go-graphql-example/cmd/gql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run -mod=mod github.com/99designs/gqlgen generate --verbose

// Resolver represents resolver type
type Resolver struct {
	todos []*model.Todo
}
