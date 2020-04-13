package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

func Upload(f *multipart.FileHeader, c *gin.Context,ch chan error) {
	//生成文件名称

	times := int(time.Now().Unix()) + rand.Int()
	split := strings.Split(f.Filename, ".")
	tmpName := strconv.Itoa(times) + "." + split[1]
	tmpPath := fmt.Sprintf("./%s/%s", "/uploads", tmpName)
	err := c.SaveUploadedFile(f, tmpPath)
	time.Sleep(time.Second * 2 )
	ch <-err

}
