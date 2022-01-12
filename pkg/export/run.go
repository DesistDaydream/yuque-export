package export

import (
	"sync"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

func Run(opts handler.YuqueUserOpts, discoveredTOCs []handler.TOC) {
	var wg sync.WaitGroup
	defer wg.Wait()

	concurrenceControl := make(chan bool, 1)

	// 逐一导出节点内容
	for _, discoveredTOC := range discoveredTOCs {
		concurrenceControl <- true
		wg.Add(1)

		go func(discoveredTOC handler.TOC) {
			defer wg.Done()
			// 获取待导出笔记的 URL
			exportURL, err := handler.GetURLForExportToc(discoveredTOC, opts)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
					"toc": discoveredTOC.Title,
					"url": exportURL,
				}).Error("获取待导出 TOC 的 URL 失败!")
			}

			// 开始导出笔记
			if opts.IsExport {
				err = ExportDoc(exportURL, discoveredTOC.Title)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"err": err,
					}).Error("导出 TOC 失败!")
				}
			}
			<-concurrenceControl
		}(discoveredTOC)

		// 介语雀不让并发太多啊。。。。。接口请求多了。。。直接限流了。。。囧
		time.Sleep(15 * time.Second)
	}
}
