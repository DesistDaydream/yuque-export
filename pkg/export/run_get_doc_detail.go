package export

import (
	"sync"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk"
	"github.com/sirupsen/logrus"
)

// 异常笔记
type ExceptionDoc struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type ExceptionDocs struct {
	ExceptionDocs []ExceptionDoc `json:"exception_docs"`
}

// 某些情况下，代替其他两个 RunXXX 函数以获取笔记详情
func GetDocDetail(h *handler.HandlerObject, tocs []yuquesdk.RepoTocData) ExceptionDocs {
	var eds ExceptionDocs

	var wg sync.WaitGroup

	concurrencyControl := make(chan bool, h.Flags.Concurrency)

	for _, toc := range tocs {
		concurrencyControl <- true

		wg.Add(1)

		go func(toc yuquesdk.RepoTocData) {
			defer wg.Done()

			// 获取 Doc 详情数据
			docDetail, err := h.Client.Doc.Get(h.Namespace, toc.Slug, &yuquesdk.DocGet{Raw: 1})
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"doc": docDetail.Data.Title,
					"err": err,
				}).Error("获取文档详情失败!")
			} else {
				var publicStatus string = ""
				if docDetail.Data.Public == 1 {
					publicStatus = "公开"
				} else {
					publicStatus = "私有"
				}
				logrus.WithFields(logrus.Fields{
					"b.文档": docDetail.Data.Title,
					"a.状态": publicStatus,
				}).Debugln("获取文档详情成功")
				// 判断文档是否为公开的。0 为私密，1 为公开
				if docDetail.Data.Public == 0 {
					eds.ExceptionDocs = append(eds.ExceptionDocs, ExceptionDoc{
						Title: docDetail.Data.Title,
						Slug:  docDetail.Data.Slug,
					})
				}
			}

			<-concurrencyControl
		}(toc)

		time.Sleep(time.Duration(h.Flags.ExportDuration) * time.Second)
	}

	wg.Wait()

	return eds
}
