package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"graphql-go-practice/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID: "TODO-3",
		Text: input.Text,
		User: &model.User{
			ID:  	input.UserID,
			Name: 	"name",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID: "TODO-1",
			Text: "todo 1",
			User: &model.User{
				ID:  	"USER-1",
				Name: 	"name",
			},
			Done: true,
		},
		{
			ID: "TODO-2",
			Text: "todo 2",
			User: &model.User{
				ID:  	"USER-2",
				Name: 	"name",
			},
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
