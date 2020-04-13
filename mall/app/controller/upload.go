package controller

import (
	"github.com/gin-gonic/gin"
	"mall/helpers"
)
type Upload struct {
}

func (u *Upload) Upload(c *gin.Context) *helpers.Response {
	// 获取上传的图片
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	ch := make(chan error,len(files))
	for _,f := range files {
		go helpers.Upload(f,c,ch)
	}
	flag := true
	var resError error
	for i:=0;i< len(files);i++ {
		if err := <-ch;err != nil {
			flag = false
			resError  = err
			break
		}
	}
	if flag {
		return  helpers.Success(nil,"上传成功")
	}else{
		return helpers.Error(400,resError.Error(),nil)
	}
}
