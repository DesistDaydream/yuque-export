package yuque

import (
	"github.com/DesistDaydream/yuque-export/pkg/handler"
)

// 实例化语雀的用户数据
func NewUserData() *UserData {
	return &UserData{}
}

// 从语雀的 API 中获取用户数据
func (u *UserData) Get(h *handler.HandlerObject, name string) error {
	endpoint := "/user"

	yc := handler.NewYuqueClient(h.Flags)
	err := yc.RequestV2("GET", endpoint, nil, u)
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
