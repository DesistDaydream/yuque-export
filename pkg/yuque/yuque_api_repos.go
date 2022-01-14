package yuque

import (
	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

func NewReposList() *ReposList {
	return &ReposList{}
}

// 从语雀的 API 中获取知识库列表
func (r *ReposList) Get(h *handler.HandlerObject, opts ...interface{}) error {
	url := YuqueBaseAPI + "/users/" + h.UserName + "/repos"
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debug("检查 URL，获取知识库列表")

	err := h.HttpHandler("GET", url, r)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReposList) Handle(h *handler.HandlerObject) error {
	panic("not implemented") // TODO: Implement
}
