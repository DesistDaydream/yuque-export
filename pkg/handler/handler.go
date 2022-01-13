package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// 用来处理语雀API的数据
type HandlerObject struct {
	// 通过 Token 获取到的用户名称
	UserName string
	// 待导出的知识库。可以是仓库的ID，也可以是以斜线分割的用户名和仓库slug的组合
	Namespace int
	// 已发现待导出的 TOCs

	// 命令行选项
	Opts YuqueUserOpts
}

// 根据命令行标志实例化一个处理器
func NewHandlerObject(opts YuqueUserOpts) *HandlerObject {
	return &HandlerObject{
		Opts: opts,
	}
}

// 处理语雀 API 是要使用的 HTTP 处理器
func (h *HandlerObject) HttpHandler(method string, url string, data YuqueData) error {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", h.Opts.Token)

	resp, err := client.Do(req)
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
