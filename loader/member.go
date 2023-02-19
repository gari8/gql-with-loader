package loader

import (
	"context"
	"github.com/gari8/gql-with-loader/entity"
	"github.com/graph-gophers/dataloader"
)

type MemberLoader struct {
	MemberRepo
}

func NewMemberLoader(memberRepo MemberRepo) *MemberLoader {
	return &MemberLoader{memberRepo}
}

func (l *MemberLoader) MembersLoaderFunc(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var ids []*string
	for _, key := range keys {
		k := key.String()
		ids = append(ids, &k)
	}

	args := ForInput(ctx)

	records, err := l.MemberRepo.FindAll(args.Status, ids)
	if err != nil {
		return []*dataloader.Result{{Error: err}}
	}

	memberByID := map[string][]*entity.Member{}
	for _, record := range records {
		memberByID[record.GroupID] = append(memberByID[record.GroupID], record)
	}

	members := make([][]*entity.Member, len(ids))
	for i, id := range ids {
		members[i] = memberByID[*id]
	}

	results := make([]*dataloader.Result, len(members))
	for i := range members {
		results[i] = &dataloader.Result{Data: members[i], Error: nil}
	}

	return results
}

type MembersInput struct {
	Status *entity.Status
}

const membersInputKey = "membersInput"

func ForInput(ctx context.Context) *MembersInput {
	return ctx.Value(membersInputKey).(*MembersInput)
}

func (l *Loaders) LoadMembers(ctx context.Context, id string, status *entity.Status) ([]*entity.Member, error) {
	loader := ctx.Value(loadersKey).(*Loaders)
	newContext := context.WithValue(ctx, membersInputKey, &MembersInput{
		Status: status,
	})
	thunk := loader.MembersByGroupID.Load(newContext, dataloader.StringKey(id))
	data, err := thunk()
	if err != nil {
		return nil, err
	}
	return data.([]*entity.Member), nil
}
