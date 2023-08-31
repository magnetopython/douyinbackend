package service

import "douyinbackend/dao"

type UserRegisterInfo struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func UserRegister(username string, password string) {
	var err error
	var info UserRegisterInfo
	var token string
	var user *dao.User

}
