package model

type ConfigData struct {
	Server struct {
		Port string `json:"port" yaml:"port"`
	} `json:"server" yaml:"server"`
	Database struct {
		Mysql struct {
			Addr     string `json:"addr" yaml:"addr"`
			Port     string `json:"port" yaml:"port"`
			Username string `json:"username" yaml:"username"`
			Password string `json:"password" yaml:"password"`
			Name     string `json:"name" yaml:"name"`
		} `json:"mysql" yaml:"mysql"`
		Redis struct {
			Addr     string `json:"addr" yaml:"addr"`
			Port     string `json:"port" yaml:"port"`
			Password string `json:"password" yaml:"password"`
		} `json:"redis" yaml:"redis"`
	} `json:"database" yaml:"database"`
	ALiYun struct {
		AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
		SignName        string `json:"signName" yaml:"signName"`
		TemplateCode    string `json:"templateCode" yaml:"templateCode"`
		EndPoint        string `json:"endPoint" yaml:"endPoint"`
		BucketName      string `json:"bucketName" yaml:"bucketName"`
	} `json:"aLiYun" yaml:"aLiYun"`
}
