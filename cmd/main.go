package main

import (
	"os"
	"sync"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/export"
	"github.com/DesistDaydream/yuque-export/pkg/get"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func LogInit(level, file string) error {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		// FieldMap:          map[logrus.fieldKey]string{},
		// CallerPrettyfier: func(*runtime.Frame) (string, string) {
		// },
		// PrettyPrint: true,
	})
	// logrus.SetFormatter(&logrus.TextFormatter{
	// 	FullTimestamp:   true,
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(l)

	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}

	return nil
}

func main() {
	// 设置命令行标志
	logLevel := pflag.String("log-level", "info", "The logging level:[debug, info, warn, error, fatal]")
	logFile := pflag.String("log-output", "", "the file which log to, default stdout")
	isExport := pflag.Bool("export", false, "是否真实导出笔记，默认不导出，仅查看可以导出的笔记")
	opts := &get.YuqueUserOpts{}
	opts.AddFlag()
	pflag.Parse()

	// 初始化日志
	if err := LogInit(*logLevel, *logFile); err != nil {
		logrus.Fatal(errors.Wrap(err, "set log level error"))
	}

	// 实例化语雀用户数据
	yud := get.NewYuqueUserData(*opts)

	// 获取待导出节点的信息
	discoveredTOCs, err := get.GetToc(yud.Opts.Token)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取待导出 TOC 信息失败！")
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	concurrenceControl := make(chan bool, 1)

	// 逐一导出节点内容
	for _, discoveredTOC := range discoveredTOCs {
		concurrenceControl <- true
		wg.Add(1)

		go func(discoveredTOC get.TOCData) {
			defer wg.Done()
			// 获取待导出笔记的 URL
			exportURL, err := yud.GetURLForExportToc(discoveredTOC)
			if err != nil {
				logrus.WithFields(logrus.Fields{"err": err, "toc": discoveredTOC.Title, "url": exportURL}).Error("获取待导出 TOC 的 URL 失败!")
			}

			// 开始导出笔记
			if *isExport {
				err = export.ExportDoc(exportURL, discoveredTOC.Title)
				if err != nil {
					logrus.WithFields(logrus.Fields{"err": err}).Error("导出 TOC 失败！")
				}
			}
			<-concurrenceControl
		}(discoveredTOC)

		// 介语雀不让并发太多啊。。。。。接口请求多了。。。直接限流了。。。囧
		time.Sleep(10 * time.Second)
	}

}
