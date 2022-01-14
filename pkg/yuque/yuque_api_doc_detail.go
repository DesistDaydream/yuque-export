package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
)

func NewDocDetail() *DocDetailData {
	return &DocDetailData{}
}

func (dd *DocDetailData) Get(h *handler.HandlerObject, name string) error {
	// 获取文档详情 URL
	endpoint := "/repos/" + fmt.Sprint(h.Namespace) + "/docs/" + name

	yc := handler.NewYuqueClient(h.Opts)
	err := yc.Request("GET", endpoint, dd)
	if err != nil {
		return err
	}

	return nil
}

func (dd *DocDetailData) Handle(h *handler.HandlerObject) error {
	return nil
}

func (dd *DocDetailData) GetDocDetailHTMLBody(h *handler.HandlerObject, slug string) (string, string, error) {
	err := dd.Get(h, slug)
	if err != nil {
		return "", "", err
	}
	return dd.Data.BodyHTML, dd.Data.Title, nil
}
