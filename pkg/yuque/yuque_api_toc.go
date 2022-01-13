package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

// TOCs 节点列表。
type TocsList struct {
	Data []TOC `json:"data"`
}

// TOCData 语雀的节点就是指一篇 笔记、表格、思维图 等等，语雀将所有内容抽象为 TOC(节点)
type TOC struct {
	// 节点类型
	Type string `json:"type"`
	// 节点名称
	Title string `json:"title"`
	// 节点唯一标识符
	UUID string `json:"uuid"`
	// 该节点的链接或 slug
	URL string `json:"url"`
	// 上一个节点的 UUID
	PrevUUID string `json:"prev_uuid"`
	// 下一个节点的 UUID
	SiblingUUID string `json:"sibling_uuid"`
	// 第一个子节点的 UUID
	ChildUUID string `json:"child_uuid"`
	// 父节点的 UUID
	ParentUUID string `json:"parent_uuid"`
	// 文档类型节点的标识符
	DocID int `json:"doc_id"`
	// 节点层级
	Level int `json:"level"`
	// 节点 ID
	ID int `json:"id"`
	// 连接是否在新窗口打开，0 当前页面打开、1 在新窗口打开
	OpenWindow int `json:"open_window"`
	// 节点是否可见。0 不可见、1 可见
	Visible int `json:"visible"`
	Depth   int `json:"depth"`
	// 节点 URL 的 slug
	Slug string `json:"slug"`
}

func NewTocsList() *TocsList {
	return &TocsList{}
}

// 从语雀的 API 中获取知识库内的文档列表
func (t *TocsList) Get(h *handler.HandlerObject) error {
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
	panic("not implemented") // TODO: Implement
}

// 根据用户设定，筛选出需要导出的文档
func (t *TocsList) DiscoveredTocs(h *handler.HandlerObject) ([]TOC, error) {
	var (
		discoveredTOCs []TOC // 已发现节点
	)

	logrus.Infof("当前知识库共有 %v 个节点", len(t.Data))

	for i := 0; i < len(t.Data); i++ {
		if t.Data[i].Depth == h.Opts.TocDepth {
			discoveredTOCs = append(discoveredTOCs, t.Data[i])
		}
	}

	// 输出一些 Debug 信息
	logrus.Infof("已发现 %v 个节点", len(discoveredTOCs))

	for _, discoveredTOC := range discoveredTOCs {
		logrus.WithFields(logrus.Fields{
			"title":         discoveredTOC.Title,
			"toc_node_uuid": discoveredTOC.UUID,
			"toc_node_url":  discoveredTOC.URL,
		}).Debug("显示已发现的节点信息")
	}

	return discoveredTOCs, nil
}
