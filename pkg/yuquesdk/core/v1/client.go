package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/utils/config"
	"github.com/sirupsen/logrus"
	// log "github.com/sirupsen/logrus"
)

var (
	// BaseAPI address
	BaseAPI = "https://www.yuque.com/api/"
	//EmptyRO empty options
	EmptyRO = new(RequestOption)
)

type Client struct {
	Token   string
	Referer string
	Cookie  string
}

type RequestOption struct {
	Method string
	Data   map[string]string
	Time   time.Duration
}

// 实例化一个向 Yuque API 发起 HTTP 请求的客户端
func NewClient(auth *config.AuthInfo) *Client {
	return &Client{
		Token:   auth.Token,
		Referer: auth.Referer,
		Cookie:  auth.Cookie,
	}
}

// 处理语雀 API 时要使用的 HTTP 处理器。现阶段只有 books/{namesapce}/export 接口会用到
func (yc *Client) Request(endpoint string, reqBody []byte, container interface{}, opts *RequestOption) (interface{}, error) {
	url := BaseAPI + endpoint
	logrus.WithFields(logrus.Fields{
		"url":     url,
		"method":  opts.Method,
		"reqBody": string(reqBody),
	}).Debug("检查发起请求时的URL")

	// 创建一个新的 Request
	req, err := http.NewRequest(opts.Method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("创建HTTP请求异常:%v", err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("referer", yc.Referer)
	req.Header.Add("cookie", yc.Cookie)
	req.Header.Add("X-Auth-Token", yc.Token)

	client := http.Client{
		Timeout: opts.Time,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("响应异常,状态:%v,错误:%v", resp.Status, err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体错误:%v", err)
	}

	err = json.Unmarshal(respBody, container)
	if err != nil {
		return nil, fmt.Errorf("解析响应体错误:%v", err)
	}

	return container, nil
}
