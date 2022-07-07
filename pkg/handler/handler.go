package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk"
	"github.com/sirupsen/logrus"
)

var (
	YuqueBaseAPI   = "https://www.yuque.com/api"
	YuqueBaseAPIV2 = "https://www.yuque.com/api/v2"
)

// 用来处理语雀API的数据
type HandlerObject struct {
	// 通过 Token 获取到的用户名称
	UserName string
	// 待导出的知识库。可以是仓库的ID，也可以是以斜线分割的用户名和仓库slug的组合
	Namespace string

	// 命令行选项
	Flags YuqueHandlerFlags

	Client *yuquesdk.Service
}

// 根据命令行标志实例化一个处理器
func NewHandlerObject(flags YuqueHandlerFlags, client *yuquesdk.Service) *HandlerObject {
	return &HandlerObject{
		Flags:  flags,
		Client: client,
	}
}

type YuqueClient struct {
	Client    *http.Client
	Token     string
	Referer   string
	Cookie    string
	UserName  string
	Namespace int
}

// 实例化一个向 Yuque API 发起 HTTP 请求的客户端
func NewYuqueClient(flags YuqueHandlerFlags) *YuqueClient {
	return &YuqueClient{
		Client:    &http.Client{Timeout: flags.Timeout},
		Token:     flags.Token,
		Referer:   flags.Referer,
		Cookie:    flags.Cookie,
		UserName:  "",
		Namespace: 0,
	}
}

// 处理语雀 API 时要使用的 HTTP 处理器。现阶段只有 books/{namesapce}/export 接口会用到
func (yc *YuqueClient) Request(method string, endpoint string, reqBody []byte, data YuqueDataHandler) error {
	url := YuqueBaseAPI + endpoint
	logrus.WithFields(logrus.Fields{
		"url":     url,
		"method":  method,
		"reqBody": string(reqBody),
	}).Debug("检查发起请求时的URL")

	// 创建一个新的 Request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("创建HTTP请求异常:%v", err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("referer", yc.Referer)
	req.Header.Add("cookie", yc.Cookie)
	req.Header.Add("X-Auth-Token", yc.Token)

	resp, err := yc.Client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("响应异常,状态:%v,错误:%v", resp.Status, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体错误:%v", err)
	}

	err = json.Unmarshal(respBody, data)
	if err != nil {
		return fmt.Errorf("解析响应体错误:%v", err)
	}

	return nil
}
