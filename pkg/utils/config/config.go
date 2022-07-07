package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AuthInfo struct {
	UserName string `json:"username"`
	Token    string `yaml:"token"`
	Cookie   string `yaml:"cookie"`
	RepoName string `yaml:"reponame"`
}

func NewAuthInfo(file string) (authinfo *AuthInfo) {
	fileByte, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(fileByte, &authinfo)
	if err != nil {
		panic(err)
	}

	return authinfo
}
