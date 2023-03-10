package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MoeSinon/paap/ent"
	"github.com/MoeSinon/paap/graph/generated"
	"github.com/MoeSinon/paap/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.AddUserInput) (*ent.User, error) {
	return r.UserSvc.Create(ctx, user)
}

func (r *queryResolver) User(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.UserSvc.Get(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
