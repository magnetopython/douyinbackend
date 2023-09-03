package service

import (
	"douyinbackend/dao"
	"douyinbackend/utils"
	"errors"
)

type UserRegisterInfo struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func UserRegister(username string, password string) (*UserRegisterInfo, error) {
	var err error
	var info UserRegisterInfo
	var token string
	var user *dao.User

	user, err = dao.GetUserInstance().QueryUserByName(username)

	if user.Id != 0 {
		err = errors.New("用户已存在")
	}
	if err != nil {
		return nil, err
	}
	user = &dao.User{
		Name:          username,
		FollowCount:   0,
		FollowerCount: 0,
		Password:      password,
	}

	err = dao.GetUserInstance().AddUser(user)

	if err != nil {
		return nil, err
	}

	token, err = utils.GenerateToken(username, user.Id)
	if err != nil {
		return nil, err
	}
	info.Token = token
	info.UserID = user.Id

	return &info, nil
}

func UserLogin(username string, password string) (*UserRegisterInfo, error) {
	var err error
	var token string
	var user *dao.User

	user, err = dao.GetUserInstance().QueryUserByName(username)
	if user.Id == 0 {
		err = errors.New("用户不存在")
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if password != user.Password {
		err = errors.New("密码错误")
		return nil, err
	}
	token, err = utils.GenerateToken(username, user.Id)
	return &UserRegisterInfo{Token: token, UserID: user.Id}, nil
}
