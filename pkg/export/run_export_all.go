package export

import (
	"sync"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v2/models"
	"github.com/sirupsen/logrus"
)

func ExportAll(h *handler.HandlerObject, tocs []*models.RepoTocData) {
	var wg sync.WaitGroup
	defer wg.Wait()

	concurrencyControl := make(chan bool, h.Flags.Concurrency)

	for _, toc := range tocs {
		// TODO: 当每个 TOC 信息中，data.child_uuid 字段不为空时，为其创建同名文件夹。以便更好分类
		// 因为 child_uuid 为空时，该文档下就没有其他文档了
		// 这个逻辑有点像递归函数，一遍一遍循环 []TOC，每次循环都处理 depth + 1 的字段，凡是 child_uuid 不为空的就创建目录，根据其 parent_uuid 决定该文档写入到哪个目录中。
		// 这种循环好么？~上千篇文档，循环好多遍，感觉很浪费，如何在循环后就去掉这些数据呢？~
		// TODO: 同时，添加可选项，是否创建文件夹以分类文档
		concurrencyControl <- true

		wg.Add(1)

		go func(toc *models.RepoTocData) {
			defer wg.Done()

			// 获取 Doc 的 HTML 格式信息
			docDetail, err := h.Client.Doc.Get(h.Namespace, toc.Slug, &models.DocGet{Raw: 1})
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
					logrus.WithFields(logrus.Fields{
						"doc": docDetail.Data.Title,
					}).Infoln("导出 MD 成功!")

					SuccessCount++
				}
			}

			<-concurrencyControl
		}(toc)

		time.Sleep(time.Duration(h.Flags.ExportDuration) * time.Second)
	}
}
