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
func (t *TocsList) Get(h *handler.HandlerObject, opts ...interface{}) error {
	url := YuqueBaseAPI + "/repos/" + fmt.Sprint(h.Namespace) + "/toc"
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debug("检查 URL，获取 TOC 数据")

	err := h.HttpHandler("GET", url, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *TocsList) Handle(h *handler.HandlerObject) error {
	// 根据用户设定，筛选出需要导出的文档
	logrus.Infof("当前知识库共有 %v 个节点", len(t.Data))
	j := 0
	for i := 0; i < len(t.Data); i++ {
		if t.Data[i].Depth == h.Opts.TocDepth {
			h.DiscoveredTocsList = append(h.DiscoveredTocsList, handler.Toc{})
			h.DiscoveredTocsList[j].Title = t.Data[i].Title
			h.DiscoveredTocsList[j].URL = t.Data[i].URL
			h.DiscoveredTocsList[j].UUID = t.Data[i].UUID
			j++
		}
	}

	// 输出一些 Debug 信息
	logrus.Infof("已发现 %v 个节点", len(h.DiscoveredTocsList))

	for _, toc := range h.DiscoveredTocsList {
		logrus.WithFields(logrus.Fields{
			"title":         toc.Title,
			"toc_node_uuid": toc.UUID,
			"toc_node_url":  toc.URL,
		}).Debug("显示已发现 TOC 的信息")
	}

	return nil
}

func (t *TocsList) GetTocsSlug(h *handler.HandlerObject) {
	logrus.Infof("当前知识库共有 %v 个节点", len(t.Data))

	for _, toc := range t.Data {
		h.DocsSlug = append(h.DocsSlug, toc.Slug)
	}
}
