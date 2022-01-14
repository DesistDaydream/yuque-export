package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	YuqueBaseAPI = "https://www.yuque.com/api/v2"
)

// 用来处理语雀API的数据
type HandlerObject struct {
	// 通过 Token 获取到的用户名称
	UserName string
	// 待导出的知识库。可以是仓库的ID，也可以是以斜线分割的用户名和仓库slug的组合
	Namespace string
	// 命令行选项
	Opts YuqueOpts
}

// 根据命令行标志实例化一个处理器
func NewHandlerObject(opts YuqueOpts) *HandlerObject {
	return &HandlerObject{
		Opts: opts,
	}
}

type YuqueClient struct {
	Client    *http.Client
	Token     string
	UserName  string
	Namespace int
}

// 实例化一个向 Yuque API 发起 HTTP 请求的客户端
func NewYuqueClient(opts YuqueOpts) *YuqueClient {
	return &YuqueClient{
		Client: &http.Client{
			Timeout: opts.Timeout,
		},
		Token:     opts.Token,
		UserName:  "",
		Namespace: 0,
	}
}

// 处理语雀 API 时要使用的 HTTP 处理器
func (yc *YuqueClient) Request(method string, endpoint string, data YuqueData) error {
	url := YuqueBaseAPI + endpoint
	logrus.WithFields(logrus.Fields{
		"url":    url,
		"method": method,
	}).Debug("检查发起请求时的URL")

	// 创建一个新的 Request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", yc.Token)

	resp, err := yc.Client.Do(req)
	if err != nil {
		logrus.Error("获取响应体错误")
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("读取响应体错误")
		return err
	}

	err = json.Unmarshal(respBody, data)
	if err != nil {
		logrus.Error("解析响应体错误")
		return err
	}

	return nil
}
