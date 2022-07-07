package service

import (
	"fmt"

	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
)

// UserService encapsulate authenticated token
type UserService struct {
	client *core.Client
}

// NewUser create User for external use
func NewUser(client *core.Client) *UserService {
	return &UserService{
		client: client,
	}
}

// Get user
func (c UserService) Get(login string) (core.UserInfo, error) {
	var (
		url  string
		user core.UserInfo
	)
	if len(login) > 0 {
		url = fmt.Sprintf("users/%s", login)
	} else {
		url = "user"
	}
	_, err := c.client.RequestObj(url, &user, core.EmptyRO)
	if err != nil {
		return user, err
	}
	return user, nil
}
