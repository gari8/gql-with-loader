package loader

import (
	"context"
	"fmt"
	"github.com/gari8/gql-with-loader/entity"
	"github.com/graph-gophers/dataloader"
)

type GroupLoader struct {
	GroupRepo
}

func NewGroupLoader(groupRepo GroupRepo) *GroupLoader {
	return &GroupLoader{groupRepo}
}

func (l *GroupLoader) GroupLoaderFunc(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var ids []string
	for _, key := range keys {
		ids = append(ids, key.String())
	}

	records, err := l.GroupRepo.FindAll(ids)
	if err != nil {
		return []*dataloader.Result{}
	}

	groupByID := map[string]*entity.Group{}
	for _, record := range records {
		groupByID[record.ID] = record
	}

	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		k := key.String()
		results[i] = &dataloader.Result{Data: nil, Error: nil}
		if group, ok := groupByID[k]; ok {
			results[i].Data = group
		} else {
			results[i].Error = fmt.Errorf("group[key=%s] not found", k)
		}
	}

	return results
}

func (l *Loaders) LoadGroup(ctx context.Context, id string) (*entity.Group, error) {
	loader := ctx.Value(loadersKey).(*Loaders)
	thunk := loader.GroupByID.Load(ctx, dataloader.StringKey(id))
	data, err := thunk()
	if err != nil {
		return nil, err
	}
	return data.(*entity.Group), nil
}
