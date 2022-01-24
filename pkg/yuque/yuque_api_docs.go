package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
)

// 实例化文档列表
func NewDocsList() *DocsList {
	return &DocsList{}
}

func (d *DocsList) Get(h *handler.HandlerObject, name string) error {
	endpoint := fmt.Sprintf("/repos/%s/docs", name)

	yc := handler.NewYuqueClient(h.Flags)
	err := yc.RequestV2("GET", endpoint, nil, d)
	if err != nil {
		return err
	}

	return nil
}

func (d *DocsList) Handle(h *handler.HandlerObject) error {
	return nil
}
