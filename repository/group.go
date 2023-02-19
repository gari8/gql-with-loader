package repository

import (
	"github.com/gari8/gql-with-loader/entity"
	"github.com/gari8/gql-with-loader/repository/model"
	"gorm.io/gorm"
)

type GroupRepo struct {
	*gorm.DB
}

func NewGroupRepo(db *gorm.DB) *GroupRepo {
	return &GroupRepo{db}
}

func (r GroupRepo) FindAll(ids []string) ([]*entity.Group, error) {
	var groups []*model.Group
	db := r.DB
	if ids != nil {
		db = db.Where("id IN ?", ids)
	}
	if err := db.Find(&groups).Error; err != nil {
		return nil, err
	}
	var result []*entity.Group
	for _, group := range groups {
		result = append(result, group.ToEntity())
	}
	return result, nil
}

func (r GroupRepo) FindByID(id string) (*entity.Group, error) {
	var group model.Group
	if err := r.DB.Where("id = ?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return group.ToEntity(), nil
}
