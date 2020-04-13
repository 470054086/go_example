package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
	var r request.IndexRequest
	err := c.BindJSON(&r)
	if err != nil {
		err := errors.Wrap(err, "");
		panic(commom.NewParamsError(err,"参数错误"))
	}
	user, err := service.S_User.AddUser(&r)
	if err != nil {
		return helpers.Success(nil,"添加失败")
	}
	return helpers.Success(request.IndexResponse{
		Id:       user,
		Mobile:   r.Mobile,
		Password: r.Password,
		Sex:      r.Sex,
	},"请求成功")
}