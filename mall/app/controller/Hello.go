package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"mall/Providers/SendMobile"
	"mall/app/commom"
	"mall/app/request"
	"mall/app/service"
	"mall/helpers"
)

type Hello struct {
}

/**

 */
func (h *Hello) Add(c *gin.Context) *helpers.Response {
	//var r request.IndexRequest
	//err := c.BindJSON(&r)
	//if err != nil {
	//	err := errors.Wrap(err, err.Error())
	//	panic(commom.NewParamsError(err, err.Error()))
	//}
	r := request.IndexRequest{
		Mobile:   "12",
		Password: "12",
		Sex:      0,
	}
	user, err := service.S_User.AddUser(&r)
	if err != nil {
		return helpers.Success(nil, "添加失败")
	}
	send := &SendMobile.SendMessage{
		Mobile: r.Mobile,
		Msg:    r.Password,
		Types:  0,
	}
	SendMobile.G_MobileSend.Send(send)
	return helpers.Success(request.IndexResponse{
		Id:       user,
		Mobile:   r.Mobile,
		Password: r.Password,
		Sex:      r.Sex,
	}, "请求成功")
	return  nil
}

func (h *Hello) Lists(c *gin.Context) *helpers.Response {
	var r request.ListRequest
	err := c.BindJSON(&r)
	if err != nil {
		err := errors.Wrap(err, err.Error())
		panic(commom.NewParamsError(err, err.Error()))
	}
	list := service.S_User.GetList(r.Mobile, r.Sex)
	return helpers.Success(list, "请求成功")
}
