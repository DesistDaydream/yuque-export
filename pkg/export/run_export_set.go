package export

import (
	"sync"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/utils/config"
	modelsv1 "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v1/models"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v2/models"
	"github.com/sirupsen/logrus"
)

var (
	SuccessCount int
	FailureCount int
)

func ExportSet(h *handler.HandlerObject, tocs []*models.RepoTocData, auth config.AuthInfo) {
	// 并发
	var wg sync.WaitGroup
	defer wg.Wait()
	// 控制并发
	concurrencyControl := make(chan bool, h.Flags.Concurrency)

	// 逐一导出节点内容
	for _, toc := range tocs {
		// 控制并发
		concurrencyControl <- true
		// 并发
		wg.Add(1)

		go func(toc *models.RepoTocData) {
			defer wg.Done()

			// 获取待导出笔记的 URL
			request := &modelsv1.BookExportRequest{
				Type:         "lakebook",
				Force:        0,
				Title:        toc.Title,
				TocNodeUUID:  toc.UUID,
				TocNodeURL:   toc.URL,
				WithChildren: true,
			}
			resp, err := h.ClientV1.BookExport.GetDownloadURL(request, h.Flags.Timeout, h.RepoID)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
					"toc": toc.Title,
				}).Error("获取待导出 TOC 的 URL 失败!")
				FailureCount++
			} else {
				logrus.WithFields(logrus.Fields{
					"toc_title":  toc.Title,
					"toc_uuid":   toc.UUID,
					"export_url": resp.Data.URL,
				}).Infof("获取待导出 TOC 的 URL 成功!")

				// 导出笔记
				if h.Flags.IsExport {
					err = ExportLakebook(resp.Data.URL, h.Flags.Path, toc.Title)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"err": err,
						}).Error("导出 TOC 失败!")
						FailureCount++
					} else {
						SuccessCount++
					}
				}
			}

			// 控制并发
			<-concurrencyControl
		}(toc)

		// 控制并发。。。。。接口请求多了。。。直接限流了。。。囧
		// 其实主要是对 GetURlForExportToc 中的接口限流，防止请求过多，导致服务器处理很多压缩任务
		time.Sleep(time.Duration(h.Flags.ExportDuration) * time.Second)
	}
}
