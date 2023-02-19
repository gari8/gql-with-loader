package loader

import (
	"context"
	"github.com/gari8/gql-with-loader/entity"
	"github.com/graph-gophers/dataloader"
	"net/http"
	"time"
)

type loadersKeyType string

const loadersKey loadersKeyType = "loaders"

type Loaders struct {
	GroupByID        Loader
	MembersByGroupID Loader
}

type Loader interface {
	ClearAll() dataloader.Interface
	Load(ctx context.Context, key dataloader.Key) dataloader.Thunk
}

type GroupRepo interface {
	FindAll(ids []string) ([]*entity.Group, error)
}

type MemberRepo interface {
	FindAll(status *entity.Status, groupIds []*string) ([]*entity.Member, error)
}

func DataLoaderMiddleware(
	loaders *Loaders,
	next http.Handler,
) http.Handler {
	loaders.GroupByID.ClearAll()
	loaders.MembersByGroupID.ClearAll()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func NewLoaders(
	groupRepo GroupRepo,
	memberRepo MemberRepo,
) *Loaders {
	groupLoader := NewGroupLoader(groupRepo)
	memberLoader := NewMemberLoader(memberRepo)
	return &Loaders{
		GroupByID:        NewBatchedFunc(groupLoader.GroupLoaderFunc),
		MembersByGroupID: NewBatchedFunc(memberLoader.MembersLoaderFunc),
	}
}

func NewBatchedFunc(bf dataloader.BatchFunc) *dataloader.Loader {
	return dataloader.NewBatchedLoader(bf, dataloader.WithWait(2*time.Millisecond))
}
