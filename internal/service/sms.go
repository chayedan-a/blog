package service

import (
	"blog/config"
	"blog/internal/model/result"
	"blog/internal/utils"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
)

const (
	SmsCodeRedisKey = "login:sms:phone:%s"
)

func (s *Service) Sms(c *gin.Context) {
	aLiYun := config.Config().ALiYun
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", aLiYun.AccessKeyId, aLiYun.AccessKeySecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	phone := c.Param("phone")
	request.PhoneNumbers = phone
	request.SignName = aLiYun.SignName
	request.TemplateCode = aLiYun.TemplateCode
	code := utils.CreateCaptcha()
	request.TemplateParam = "{\"code\":\"" + code + "\"}"
	_, err = client.SendSms(request)
	if err != nil {
		//todo 错误日志
		result.Fail(c, result.Error)
		return
	}
	s.dao.Redis.Do("SET", fmt.Sprintf(SmsCodeRedisKey, phone), code, "EX", 5*60)
	result.Ok(c, "发送成功")
}
