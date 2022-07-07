package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 认证信息配置
type AuthConfig struct {
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
	Cookie   string `yaml:"cookie"`
	Referer  string `yaml:"referer"`
}

func NewAuthInfo(file string) (auth *AuthConfig) {
	// 读取认证信息
	fileByte, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(fileByte, &auth)
	if err != nil {
		panic(err)
	}
	return auth
}

// // 判断文件中是否存在域名
// func (c *AuthConfig) IsUserExist(userName string) bool {
// 	if _, ok := c.AuthList[userName]; ok {
// 		return true
// 	} else {
// 		return false
// 	}
// }
