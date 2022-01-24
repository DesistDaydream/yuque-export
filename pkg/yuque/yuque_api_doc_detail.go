package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
)

// 实例化文档详情
func NewDocDetail() *DocDetailData {
	return &DocDetailData{}
}

func (dd *DocDetailData) Get(h *handler.HandlerObject, name string) error {
	// 获取文档详情 URL
	endpoint := fmt.Sprintf("/repos/%s/docs/%s", h.Namespace, name)

	yc := handler.NewYuqueClient(h.Flags)
	err := yc.RequestV2("GET", endpoint, nil, dd)
	if err != nil {
		return err
	}

	return nil
}

func (dd *DocDetailData) Handle(h *handler.HandlerObject) error {
	return nil
}
