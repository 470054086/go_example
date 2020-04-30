package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"mall/Providers/SendMobile"
	"mall/Providers/Subscribe"
	"mall/app/commom"
	"mall/app/request"
	"mall/app/service"
	"mall/helpers"
)

type Hello struct {
}

func (h *Hello) Test(c *gin.Context) *helpers.Response {
	Subscribe.G_SUB.AddQueue("shop", "pay3", 2)
	Subscribe.G_SUB.AddQueue("shop", "pay3", 3)
	Subscribe.G_SUB.SendQueue("shop", "pay3", func(i int) {
		fmt.Println(i)
	})
	return nil
}

/**

 */
func (h *Hello) Add(c *gin.Context) *helpers.Response {
	var r request.IndexRequest
	err := c.ShouldBind(&r)
	if err != nil {
		err := errors.Wrap(err, err.Error())
		panic(commom.NewParamsError(err, err.Error()))
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
}

// 查询列表
func (h *Hello) Lists(c *gin.Context) *helpers.Response {
	var r request.ListRequest
	err := c.ShouldBind(&r)
	if err != nil {
		err := errors.Wrap(err, err.Error())
		panic(commom.NewParamsError(err, err.Error()))
	}
	list := service.S_User.GetList(r.Mobile, r.Sex)
	return helpers.Success(list, "请求成功")
}

func (h *Hello) Login(c *gin.Context) *helpers.Response {
	var r request.LoginRequest
	err := c.ShouldBind(&r)
	if err != nil {
		err := errors.Wrap(err, err.Error())
		panic(commom.NewParamsError(err, err.Error()))
	}
	service.S_User.Login(r.Mobile, r.Password)
	s := make(map[string]string)
	return helpers.Success(s, "请求成功")
}
