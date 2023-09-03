package dao

import (
	"douyinbackend/config"
	"sync"
)

type User struct {
	Id            int64  `gorm:"column:user_id" json:"id"`
	Name          string `gorm:"column:user_name" json:"name" `
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"`
	IsFollow      bool   `gorm:"column:is_follow" json:"is_follow"`
	Password      string `gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}

var user *User

var userOnce sync.Once

func GetUserInstance() *User {
	userOnce.Do(func() {
		user = &User{}
	})
	return user
}

func (User) AddUser(user *User) error {
	tx := config.Db.Begin()
	db := tx.Create(user)
	err := db.Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (User) QueryUserByName(username string) (*User, error) {
	var user User
	tx := config.Db.Begin()
	db := tx.Where("user_name=?", username).Find(&user)
	err := db.Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

func (User) QueryUserById(userId int64) (*User, error) {
	var user User
	tx := config.Db.Begin()
	db := tx.Where("user_id=?", userId).Find(&user)
	err := db.Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}
