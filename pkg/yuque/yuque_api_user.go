package yuque

import (
	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

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
	return nil
}

func (u *UserData) GetUserName() string {
	return u.Data.Name
}
