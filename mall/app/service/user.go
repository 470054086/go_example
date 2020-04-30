package service

import (
	"mall/app/Inc"
	"mall/app/commom"
	"mall/app/constaent"
	"mall/app/model"
	"mall/app/request"
	"mall/app/tool"
	"sync"
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
		Password: tool.PasswordHash(r.Password),
		Sex:      r.Sex,
	}
	return userModel.Create(userModel)
}

func (u *User) GetList(mobile string, sex int) *request.ListResponse {
	list := u.userModel.GetList(mobile, sex)
	var sy sync.WaitGroup
	sy.Add(2)
	var result []*request.UserListResponse
	var count int
	go func() {
		result = make([]*request.UserListResponse, len(list))
		for k, v := range list {
			response := &request.UserListResponse{
				Mobile:  v.Mobile,
				Sex:     v.Sex,
				SexName: Inc.SexType(v.Sex).String(),
			}
			result[k] = response
		}
		sy.Done()
	}()
	go func() {
		count = u.userModel.GetListCunt(mobile, sex)
		sy.Done()
	}()
	sy.Wait()
	return &request.ListResponse{
		List:     result,
		Total:    count,
		Page:     1,
		PageSize: 1,
	}
}

func (u *User) Login(mobile, password string) bool {
	user := u.userModel.GetFirst(mobile)
	if user == nil {
		panic(commom.NewParamsErrorCode(nil, constaent.UserNotExistError.String(), int(constaent.UserNotExistError)))
	}
	if verify := tool.PasswordVerify(password, user.Password); !verify {
		panic(commom.NewParamsErrorCode(nil, constaent.UserPasswordError.String(), int(constaent.UserPasswordError)))
	}
	return true
}
