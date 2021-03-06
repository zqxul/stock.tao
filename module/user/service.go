package user

import (
	"time"

	"stock.tao/dao"
	"stock.tao/util"
)

func init() {

}

// UsernameExist ==> check if the username already exists
func UsernameExist(username string) bool {
	userCondition := dao.UserCondition{Username: &username}
	return dao.Exist(&userCondition)
}

// CreateUser ==> create new user
func CreateUser(username, password, email, nickname string) uint64 {
	salt := util.Salt(32)
	encPwd := util.Encrypt([]byte(salt), []byte(password))
	user := &dao.User{
		ID:         util.NextID(),
		Username:   username,
		Password:   encPwd,
		Salt:       salt,
		Email:      email,
		Nickname:   nickname,
		CreateTime: time.Now(),
		UpdateTime: time.Time{},
		Delete:     false,
	}
	return dao.InsertUser(user)
}

// VerifyUser ==> verify user
func VerifyUser(username, password string) (bool, *dao.User) {
	userCondition := dao.UserCondition{
		Username: &username,
	}
	user := dao.SelectOne(&userCondition)
	if user == nil {
		return false, nil
	}
	return user.Password == util.Encrypt([]byte(user.Salt), []byte(password)), user
}
