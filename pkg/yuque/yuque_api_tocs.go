package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

func NewTocsList() *TocsList {
	return &TocsList{}
}

// 从语雀的 API 中获取知识库内的文档列表
func (t *TocsList) Get(h *handler.HandlerObject, name string) error {
	fmt.Println(name)
	endpoint := "/repos/" + h.Namespace + "/toc"

	yc := handler.NewYuqueClient(h.Opts)
	err := yc.Request("GET", endpoint, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *TocsList) Handle(h *handler.HandlerObject) error {
	return nil
}

func (t *TocsList) DiscoverTocs(h *handler.HandlerObject) []TOC {
	var discoveredTocs []TOC
	// 根据用户设定，筛选出需要导出的文档
	logrus.Infof("当前知识库共有 %v 个节点", len(t.Data))
	for _, data := range t.Data {
		if data.Depth == h.Opts.TocDepth {
			discoveredTocs = append(discoveredTocs, data)

		}
	}

	return discoveredTocs
}

func (t *TocsList) GetTocsSlug(h *handler.HandlerObject) {
	logrus.Infof("当前知识库共有 %v 个节点", len(t.Data))

	for _, toc := range t.Data {
		h.DocsSlug = append(h.DocsSlug, toc.Slug)
	}
}
