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

// 是否有必要抽象出来这层接口？
type YuqueData interface {
	Get(handler *HandlerObject) error
}

// 处理语雀 API 是要使用的 HTTP 处理器
func HttpHandler(method, url, token string, data YuqueData) error {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", token)

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
