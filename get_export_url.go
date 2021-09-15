package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

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

// GetURLForExportToc is
func GetURLForExportToc(tocdata TOCData, cookie string) (string, error) {
	url := "https://www.yuque.com/api/books/11199981/export"
	method := "POST"

	// 根据节点信息，配置当前待导出节点的请求体信息
	reqBodyForExportToc := ReqBodyForExportToc{
		Type:         "lakebook",
		Force:        0,
		Title:        tocdata.Title,
		TocNodeUUID:  tocdata.UUID,
		TocNodeURL:   tocdata.URL,
		WithChildren: true,
	}

	// 解析请求体
	reqBodyByte, err := json.Marshal(reqBodyForExportToc)
	if err != nil {
		return "", err
	}
	logrus.Debug("待导出笔记信息的解析结果：", string(reqBodyByte))

	// 实例化 HTTP 请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		return "", err
	}

	req.Header.Add("authority", "www.yuque.com")
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-csrf-token", "3t1WleUAVuVSM5P-82xwT3Bl")
	// req.Header.Add("x-requested-with", "XMLHttpRequest")
	// req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://www.yuque.com")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://www.yuque.com/desistdaydream/learning")
	req.Header.Add("cookie", cookie)
	req.Header.Add("Content-Type", "application/json")

	// 建立连接
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	logrus.Debug("获取导出 URL 时的响应码：", resp.Status)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var exportDatas ExportDatas

	err = json.Unmarshal(respBody, &exportDatas)
	if err != nil {
		return "", err
	}

	return exportDatas.Data.URL, nil
}

type ExportDatas struct {
	Data ExportData `json:"data"`
}

type ExportData struct {
	State string `json:"state"`
	URL   string `json:"url"`
}
