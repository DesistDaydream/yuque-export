package export

import (
	"sync"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/yuque"
	"github.com/sirupsen/logrus"
)

func Run(h *handler.HandlerObject) {
	var wg sync.WaitGroup
	defer wg.Wait()

	concurrenceControl := make(chan bool, 1)

	// 逐一导出节点内容
	for _, tocsInfo := range h.DiscoveredTocsList {
		concurrenceControl <- true
		wg.Add(1)

		go func(tocsInfo handler.Toc) {
			defer wg.Done()
			// 获取待导出笔记的 URL
			exportURL, err := yuque.GetURLForExportToc(h, tocsInfo)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
					"toc": tocsInfo,
					"url": exportURL,
				}).Error("获取待导出 TOC 的 URL 失败!")
			}

			// 开始导出笔记
			if h.Opts.IsExport {
				err = ExportDoc(exportURL, tocsInfo.Title)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"err": err,
					}).Error("导出 TOC 失败!")
				}
			}
			<-concurrenceControl
		}(tocsInfo)

		// 介语雀不让并发太多啊。。。。。接口请求多了。。。直接限流了。。。囧
		// 其实主要是对 GetURlForExportToc 中的接口限流，防止请求过多，导致服务器处理很多压缩任务
		time.Sleep(time.Duration(h.Opts.ExportDuration) * time.Second)
	}
}
