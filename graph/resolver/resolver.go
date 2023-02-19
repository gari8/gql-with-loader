package resolver

import (
	"context"
	"github.com/gari8/gql-with-loader/entity"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type GroupRepo interface {
	FindAll(ids []string) ([]*entity.Group, error)
	FindByID(id string) (*entity.Group, error)
}
type MemberRepo interface {
	FindAll(status *entity.Status, groupIds []*string) ([]*entity.Member, error)
	FindByID(id string) (*entity.Member, error)
}
type Loaders interface {
	LoadGroup(ctx context.Context, id string) (*entity.Group, error)
	LoadMembers(ctx context.Context, id string, status *entity.Status) ([]*entity.Member, error)
}

type Resolver struct {
	GroupRepo
	MemberRepo
	Loaders
}

func NewResolver(groupRepo GroupRepo, memberRepo MemberRepo, loaders Loaders) *Resolver {
	return &Resolver{groupRepo, memberRepo, loaders}
}
