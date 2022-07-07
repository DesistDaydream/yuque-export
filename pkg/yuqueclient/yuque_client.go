package yuqueclient

import "net/http"

type YuqueClient struct {
	goHttpClient *http.Client
}

func NewYuqueClient() *YuqueClient {
	return &YuqueClient{
		goHttpClient: &http.Client{},
	}
}

// TODO: 我想自己设计一下语雀的 SDK~~~
