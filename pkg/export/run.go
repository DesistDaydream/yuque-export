package export

import (
	"sync"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/yuque"
	"github.com/sirupsen/logrus"
)

func RunSet(h *handler.HandlerObject, tocs []yuque.TOC) {
	var wg sync.WaitGroup
	defer wg.Wait()

	// 介语雀不让并发太多啊
	concurrenceControl := make(chan bool, 1)

	// 逐一导出节点内容
	for _, toc := range tocs {
		// 介语雀不让并发太多啊
		concurrenceControl <- true

		wg.Add(1)

		go func(toc yuque.TOC) {
			defer wg.Done()

			// 获取待导出笔记的 URL
			exportURL, err := yuque.GetURLForExportToc(h, toc)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
					"toc": toc.Title,
				}).Error("获取待导出 TOC 的 URL 失败!")
			}

			// 开始导出笔记
			if h.Flags.IsExport {
				err = ExportDoc(exportURL, toc.Title)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"err": err,
					}).Error("导出 TOC 失败!")
				}
			}

			// 介语雀不让并发太多啊
			<-concurrenceControl
		}(toc)

		// 介语雀不让并发太多啊。。。。。接口请求多了。。。直接限流了。。。囧
		// 其实主要是对 GetURlForExportToc 中的接口限流，防止请求过多，导致服务器处理很多压缩任务
		time.Sleep(time.Duration(h.Flags.ExportDuration) * time.Second)
	}
}

func RunOne(h *handler.HandlerObject, tocs []yuque.TOC) {
	var wg sync.WaitGroup
	defer wg.Wait()

	concurrenceControl := make(chan bool, 5)

	for _, toc := range tocs {
		concurrenceControl <- true

		wg.Add(1)

		slug := toc.Slug

		go func(slug string) {
			defer wg.Done()

			// 获取 Doc 详情
			docDetail := yuque.NewDocDetail()

			// 获取文档的 HTML 格式信息
			body, name, err := docDetail.GetDocDetailBodyHTML(h, slug)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
					"toc": name,
				}).Error("获取待导出 DOC 失败!")
			}

			// 导出笔记
			err = ExportMd(body, name)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
				}).Error("导出 MD 失败!")
			}

			<-concurrenceControl
		}(slug)
	}
}
