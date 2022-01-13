package yuque

import (
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

// 用户数据
type UserData struct {
	Data User `json:"data"`
}

type User struct {
	ID             int       `json:"id"`
	Type           string    `json:"type"`
	Login          string    `json:"login"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	AvatarURL      string    `json:"avatar_url"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Serializer     string    `json:"_serializer"`
}

// 实例化语雀的用户数据
func NewUserData() *UserData {
	return &UserData{}
}

// 从语雀的 API 中获取用户数据
func (u *UserData) Get(h *handler.HandlerObject, opts ...interface{}) error {
	url := YuqueBaseAPI + "/user"
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debug("检查 URL，获取用户数据")

	err := h.HttpHandler("GET", url, u)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserData) Handle(h *handler.HandlerObject) error {
	// 将获取用户名称赋值给 handler 中的 UserName
	h.UserName = u.Data.Name
	return nil
}
