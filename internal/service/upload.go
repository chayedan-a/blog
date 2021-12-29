package service

import (
	"fmt"
	"os"

	"blog/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

func (s *Service) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	//allowExts := []string {".jpg", ".png", ".gif", ".jpeg", ".doc", ".docx", ".ppt", ".pptx", ".xls", ".xlsx", ".pdf"}
	aLiYun := config.Config().ALiYun
	// 创建OSSClient实例。
	client, err := oss.New(aLiYun.EndPoint, aLiYun.AccessKeyId, aLiYun.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(aLiYun.BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fd, err := file.Open()
	// 读取本地文件。

	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject("<yourObjectName>", fd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
