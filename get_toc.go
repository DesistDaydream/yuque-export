package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// TOCs 节点信息。
type TOCs struct {
	Data []TOCData `json:"data"`
}

// GetToc is 获取指定 namespace 下的 Toc 目录
func GetToc(token string) ([]TOCData, error) {
	// TODO 提出来，当做用户信息，写到一个 struct 中
	namespace := "desistdaydream/learning"

	url := "https://www.yuque.com/api/v2/repos/" + namespace + "/toc"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var (
		tocs           TOCs      // 节点信息
		discoveredTOCs []TOCData // 已发现节点
	)

	json.Unmarshal(respBody, &tocs)

	logrus.Infof("当前知识库共有 %v 个节点", len(tocs.Data))

	for i := 0; i < len(tocs.Data); i++ {
		if tocs.Data[i].Depth == 2 && tocs.Data[i].ParentUUID == "KXnLUubs9BNnzrs0" {
			discoveredTOCs = append(discoveredTOCs, tocs.Data[i])
		}
	}

	// 输出一些 Debug 信息
	logrus.Infof("已发现 %v 个节点", len(discoveredTOCs))

	for _, discoveredTOC := range discoveredTOCs {
		// logrus.WithFields(logrus.Fields{"toc": discoveredTOC.Title}).Debugf("显示已发现的节点")
		logrus.WithFields(logrus.Fields{
			"title":         discoveredTOC.Title,
			"toc_node_uuid": discoveredTOC.UUID,
			"toc_node_url":  discoveredTOC.URL,
		}).Debug("显示已发现的节点信息")
	}

	return discoveredTOCs, nil
}

// TOCData 语雀的节点就是指一篇 笔记、表格、思维图 等等，语雀将所有内容抽象为 TOC(节点)
type TOCData struct {
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
