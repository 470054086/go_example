package service

import (
	"mall/app/model"
	"mall/app/request"
)

type User struct {
	userModel *model.User
}

var S_User *User

func init() {
	S_User = &User{userModel: &model.User{}}
}

func (u *User) AddUser(r *request.IndexRequest) (uint, error) {
	userModel := &model.User{
		Mobile:   r.Mobile,
		Password: r.Password,
		Sex:      r.Sex,
	}
	return userModel.Create(userModel)
}

func (u *User) GetList(mobile string, sex int) []*model.User {
	return u.userModel.GetList(mobile, sex)
}
