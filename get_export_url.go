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
func GetURLForExportToc(tocdata TOCData) (string, error) {
	// 根据节点信息，配置当前待导出节点的请求体信息
	reqBodyForExportToc := ReqBodyForExportToc{
		Type:         "lakebook",
		Force:        0,
		Title:        tocdata.Title,
		TocNodeUUID:  tocdata.UUID,
		TocNodeURL:   tocdata.URL,
		WithChildren: true,
	}

	url := "https://www.yuque.com/api/books/11199981/export"
	method := "POST"
	cookie := "lang=zh-cn; UM_distinctid=179878f14125d7-03c4269457230e-d7e1938-13c680-179878f1413741; _yuque_session=sLIjjTe_9QQlnNpb8Yqp3JmXGeN5fAZ7RUg0_yio1-eMSxOQkYsDLWCta4OHciN_B3HxLo6UjSrwG0V6paVIlw==; yuque_ctoken=3t1WleUAVuVSM5P-82xwT3Bl; _TRACERT_COOKIE__SESSION=b953f580-10e0-425d-9f54-d2367dfeeec4; CNZZDATA1272061571=342748924-1621473690-%7C1627262936; acw_tc=0bda731d16272645601475651e94cbe981689b1cba5086f935f95bd9a2b734; tree=a385%01fe6f8647-d16e-433d-a56e-cc00418220fd%0129"

	reqBodyByte, err := json.Marshal(reqBodyForExportToc)
	if err != nil {
		return "", err
	}
	logrus.Debug("待导出笔记信息的解析结果：", string(reqBodyByte))

	client := &http.Client{}
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
