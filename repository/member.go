package repository

import (
	"github.com/gari8/gql-with-loader/entity"
	"github.com/gari8/gql-with-loader/repository/model"
	"gorm.io/gorm"
)

type MemberRepo struct {
	*gorm.DB
}

func NewMemberRepo(db *gorm.DB) *MemberRepo {
	return &MemberRepo{db}
}

func (r MemberRepo) FindAll(status *entity.Status, groupIds []*string) ([]*entity.Member, error) {
	var members []*model.Member
	db := r.DB
	if status != nil {
		db = db.Where("status = ?", status.Atoi())
	}
	if groupIds != nil {
		db = db.Where("group_id IN ?", groupIds)
	}
	if err := db.Find(&members).Error; err != nil {
		return nil, err
	}
	var result []*entity.Member
	for _, member := range members {
		result = append(result, member.ToEntity())
	}
	return result, nil
}

func (r MemberRepo) FindByID(id string) (*entity.Member, error) {
	var member model.Member
	if err := r.DB.Where("id = ?", id).First(&member).Error; err != nil {
		return nil, err
	}
	return member.ToEntity(), nil
}
