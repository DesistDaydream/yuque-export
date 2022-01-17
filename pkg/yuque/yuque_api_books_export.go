package yuque

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

// ReqBodyForGetExportURL is
type ReqBodyForExportToc struct {
	Type         string `json:"type"`
	Force        int    `json:"force"`
	Title        string `json:"title"`
	TocNodeUUID  string `json:"toc_node_uuid"`
	TocNodeURL   string `json:"toc_node_url"`
	WithChildren bool   `json:"with_children"`
}

// GetURLForExportToc 获取待导出 TOC 的 URL
func GetURLForExportToc(h *handler.HandlerObject, toc TOC) (string, error) {
	url := "https://www.yuque.com/api/books/" + fmt.Sprint(h.Namespace) + "/export"
	method := "POST"

	// 根据节点信息，配置当前待导出节点的请求体信息
	reqBodyForExportToc := ReqBodyForExportToc{
		Type:         "lakebook",
		Force:        0,
		Title:        toc.Title,
		TocNodeUUID:  toc.UUID,
		TocNodeURL:   toc.URL,
		WithChildren: true,
	}

	// 解析请求体
	reqBodyByte, err := json.Marshal(reqBodyForExportToc)
	if err != nil {
		return "", err
	}

	// 实例化 HTTP 请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		return "", err
	}

	req.Header.Add("authority", "www.yuque.com")
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("referer", h.Flags.Referer)
	req.Header.Add("cookie", h.Flags.Cookie)

	// 建立连接
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var exportDatas ExportsData

	err = json.Unmarshal(respBody, &exportDatas)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 200 {
		logrus.WithFields(logrus.Fields{
			"toc":    toc.Title,
			"status": resp.Status,
			"url":    exportDatas.Data.URL,
		}).Infof("获取待导出 TOC 的 URL 成功!")
	} else {
		return "", fmt.Errorf(resp.Status)
	}

	return exportDatas.Data.URL, nil
}
