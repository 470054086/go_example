package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type UserInfo struct {
	User string `binding:"required"`
	Name string `binding:"required,email"`
	Age  int
}

type RequestPost struct {
	Name     string `binding:"required"`
	Age      int    `binding:"required,email"`
	RealName string `json:"real_name"`
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func (r *RequestPost) GetError(err validator.ValidationErrors) string {

	// 这里的 "LoginRequest.Mobile" 索引对应的是模型的名称和字段
	fmt.Println(err)
	return ""
}

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Abort()
		r:= map[string]interface{}{
			"name":1,
			"age":"xiaobnai",
		}
		c.JSON(400,r)
		return
		//fmt.Println("我就是个傻逼中间件")
		//c.Next()
	})
	r.GET("/user/:name", func(c *gin.Context) {
		param := c.Param("name")
		age := c.DefaultQuery("age", "20")
		atoi, _ := strconv.Atoi(age)
		c.JSON(200,UserInfo{
			User: "xiaobai",
			Name: param,
			Age : atoi,
		})
	})

	// json数据绑定
	r.POST("/post/user", func(c *gin.Context) {
		var r RequestPost
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}else{
			c.JSON(200,login)
		}
	})

	r.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		fmt.Println(files)
		for _, file := range files {
			// Upload the file to specific dst.
			err := c.SaveUploadedFile(file, "./upload/"+file.Filename)
			fmt.Println(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	r.Run(":8082")


}
