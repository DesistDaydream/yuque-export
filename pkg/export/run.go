package export

import (
	"sync"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/yuque"
	"github.com/sirupsen/logrus"
)

var (
	SuccessCount int
	FailureCount int
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
			exportsData := yuque.NewExportsData()
			err := exportsData.Get(h, toc.UUID)
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
					"export_url": exportsData.Data.URL,
				}).Infof("获取待导出 TOC 的 URL 成功!")

				// 导出笔记
				if h.Flags.IsExport {
					err = ExportDoc(exportsData.Data.URL, h.Flags.Path, toc.Title)
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

	concurrenceControl := make(chan bool, 3)

	for _, toc := range tocs {
		// TODO: 当每个 TOC 信息中，data.child_uuid 字段不为空时，为其创建同名文件夹。以便更好分类
		// 因为 child_uuid 为空时，该文档下就没有其他文档了
		// 这个逻辑有点像递归函数，一遍一遍循环 []TOC，每次循环都处理 depth + 1 的字段，凡是 child_uuid 不为空的就创建目录，根据其 parent_uuid 决定该文档写入到哪个目录中。
		// 这种循环好么？~上千篇文档，循环好多遍，感觉很浪费，如何在循环后就去掉这些数据呢？~
		// TODO: 同时，添加可选项，是否创建文件夹以分类文档
		concurrenceControl <- true

		wg.Add(1)

		go func(toc yuque.TOC) {
			defer wg.Done()

			// 获取 Doc 的 HTML 格式信息
			docDetail := yuque.NewDocDetail()
			err := docDetail.Get(h, toc.Slug)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"doc": docDetail.Data.Title,
					"err": err,
				}).Error("获取文档详情失败!")
				FailureCount++
			} else {
				// 导出笔记
				err = ExportMd(docDetail, h.Flags.Path)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"doc": docDetail.Data.Title,
						"err": err,
					}).Error("导出 MD 失败!")
					FailureCount++
				} else {
					SuccessCount++
				}
			}

			<-concurrenceControl
		}(toc)

		time.Sleep(time.Duration(h.Flags.ExportDuration) * time.Second)
	}
}
