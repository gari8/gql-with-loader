package main

import (
	"fmt"
	"github.com/dgryski/trifles/uuid"
	"github.com/gari8/gql-with-loader/repository/model"
	"github.com/mattn/go-gimei"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func NewDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(
		&model.Group{},
		&model.Member{},
	); err != nil {
		return nil, err
	}
	db = db.Debug()
	if IsEmpty(db) {
		Seed(db)
	}
	return db, nil
}

func IsEmpty(db *gorm.DB) bool {
	var cnt int64
	db.Model(&model.Group{}).Count(&cnt)
	return cnt == 0
}

func Seed(db *gorm.DB) {
	groups := createGroups(5)
	var members []*model.Member
	for _, g := range groups {
		members = append(members, createMembers(10, g.ID)...)
	}
	db.Create(groups)
	db.Create(members)
}

func createMembers(count int, groupId string) []*model.Member {
	var members []*model.Member
	for i := 1; i < count+1; i++ {
		rand.Seed(time.Now().UnixNano())
		members = append(members, &model.Member{
			ID:      uuid.UUIDv4(),
			Name:    gimei.NewName().Kanji(),
			Status:  int8(rand.Intn(3)),
			GroupID: groupId,
		})
	}
	return members
}

func createGroups(count int) []*model.Group {
	var groups []*model.Group
	for i := 1; i < count+1; i++ {
		desc := fmt.Sprintf("Group %d Description", i)
		groups = append(groups, &model.Group{
			ID:          uuid.UUIDv4(),
			Name:        fmt.Sprintf("Group %d", i),
			Description: &desc,
		})
	}
	return groups
}
