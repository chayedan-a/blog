package service

import (
	"fmt"
	"path"
	"path/filepath"
	"time"

	"blog/internal/model/result"

	"github.com/gin-gonic/gin"
)

func (s *Service) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("Upload FormFile err：%+v", err)
		result.Fail(c, result.Error)
		return
	}

	allowExts := []string{".jpg", ".png", ".jpeg"}
	ext := path.Ext(file.Filename)
	if !containsInSlice(allowExts, ext) {
		result.Fail(c, result.NotSupportedFileExt)
		return
	}
	fileHeader, err := file.Open()
	if err != nil {
		fmt.Printf("Upload Open err：%+v", err)
		result.Fail(c, result.Error)
		return
	}

	defer fileHeader.Close()

	now := time.Now()
	fileDir := fmt.Sprintf("articles/%s", now.Format("200601"))
	timeStamp := now.Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, file.Filename)
	objectName := filepath.Join(fileDir, fileName)
	fmt.Println(objectName)
	// 上传文件流。
	url, err := s.dao.OssUpload(objectName, fileHeader)
	if err != nil {
		result.Fail(c, result.Error)
		return
	}
	result.Ok(c, map[string]string{
		"path": url,
	})
}

func containsInSlice(items []string, item string) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
