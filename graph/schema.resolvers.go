package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql_go/database"
	"graphql_go/graph/generated"
	"graphql_go/graph/model"
)

var db = database.Connect()

// CreateDog is the resolver for the createDog field.
func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	return db.Save(input), nil
}

// Dog is the resolver for the dog field.
func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	return db.FindById(id), nil
}

// Dogs is the resolver for the dogs field.
func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
