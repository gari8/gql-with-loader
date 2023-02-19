package model

import (
	"github.com/gari8/gql-with-loader/entity"
	"gorm.io/gorm"
)

type Group struct {
	ID          string  `gorm:"primaryKey;varchar(36);not null;uniqueIndex;"`
	Name        string  `json:"name" gorm:"name;not null"`
	Description *string `json:"description" gorm:"description"`
	gorm.Model
}

func (g *Group) ToEntity() *entity.Group {
	return &entity.Group{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
	}
}
