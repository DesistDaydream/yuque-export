package v2

import (
	"fmt"

	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v2/models"
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
func (c UserService) Get(login string) (models.UserInfo, error) {
	var (
		url  string
		user models.UserInfo
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
