package dao

import (
	"fmt"
	"io"

	"blog/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func (d *Dao) OssUpload(objectName string, fileHeader io.Reader) (url string, err error) {
	aLiYun := config.Config().ALiYun
	// 创建OSSClient实例。
	client, err := oss.New(aLiYun.EndPoint, aLiYun.AccessKeyId, aLiYun.AccessKeySecret)
	if err != nil {
		fmt.Printf("OssUpload oss.New err：%+v", err)
		return "", err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(aLiYun.BucketName)
	if err != nil {
		fmt.Printf("OssUpload client.Bucket err：%+v", err)
		return "", err
	}

	err = bucket.PutObject(objectName, fileHeader)
	if err != nil {
		fmt.Printf("OssUpload bucket.PutObject err：%+v", err)
		return "", err
	}
	return fmt.Sprintf("https://%s.%s", aLiYun.BucketName, aLiYun.EndPoint) + objectName, nil
}
