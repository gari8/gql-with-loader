package model

import (
	"github.com/gari8/gql-with-loader/entity"
	"gorm.io/gorm"
)

type Member struct {
	ID      string `gorm:"primaryKey;varchar(36);not null;uniqueIndex;"`
	Name    string `json:"name" gorm:"name;not null"`
	Status  int8   `json:"status" gorm:"status;not null"`
	GroupID string `json:"groupId" gorm:"group_id;not null;references:groups(id)"`
	gorm.Model
}

func (m *Member) ToEntity() *entity.Member {
	return &entity.Member{
		ID:      m.ID,
		Name:    m.Name,
		Status:  entity.AllStatus[m.Status],
		GroupID: m.GroupID,
	}
}
