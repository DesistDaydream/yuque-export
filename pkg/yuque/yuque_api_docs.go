package yuque

import (
	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

func NewDocsList() *DocsList {
	return &DocsList{}
}

func (d *DocsList) Get(h *handler.HandlerObject, name string) error {
	endpoint := "/repos/" + name + "/docs"

	yc := handler.NewYuqueClient(h.Opts)
	err := yc.Request("GET", endpoint, d)
	if err != nil {
		return err
	}

	return nil
}
func (d *DocsList) Handle(h *handler.HandlerObject) error {
	logrus.Infof("当前知识库共有 %v 篇文档", len(d.Data))
	for _, doc := range d.Data {
		h.DocsSlug = append(h.DocsSlug, doc.Slug)
	}
	return nil
}
