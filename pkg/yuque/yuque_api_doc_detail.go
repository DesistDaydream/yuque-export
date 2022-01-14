package yuque

import (
	"fmt"
	"os"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/utils/converter"
	"github.com/sirupsen/logrus"
)

func NewDocDetail() *DocDetailData {
	return &DocDetailData{}
}

func (dd *DocDetailData) Get(h *handler.HandlerObject, opts ...interface{}) error {
	var doc string

	for _, opt := range opts {
		if docDetail, ok := opt.(string); ok {
			doc = docDetail
		}
	}

	// 获取文档详情 URL
	url := YuqueBaseAPI + "/repos/" + fmt.Sprint(h.Namespace) + "/docs/" + doc

	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debugf("检查 URL，获取 %v 仓库中,文档 %v 的详情", h.Opts.RepoName, doc)

	err := h.HttpHandler("GET", url, dd)
	if err != nil {
		return err
	}

	return nil
}

func (dd *DocDetailData) Handle(h *handler.HandlerObject) error {
	mark, err := converter.ConvertHTML2Markdown(dd.Data.BodyHTML)
	if err != nil {
		return err
	}

	b := []byte(mark)
	fileName := "./files/" + dd.Data.Title + ".md"
	os.WriteFile(fileName, b, 0666)

	return nil
}

func (dd *DocDetailData) GetDocDetailBody() {

}
