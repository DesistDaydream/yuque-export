package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

func NewDocsList() *DocsList {
	return &DocsList{}
}

func (d *DocsList) Get(h *handler.HandlerObject, opts ...interface{}) error {
	url := YuqueBaseAPI + "/repos/" + fmt.Sprint(h.Namespace) + "/docs"
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debugf("检查 URL，获取%v仓库的文档列表", h.Opts.RepoName)

	err := h.HttpHandler("GET", url, d)
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
