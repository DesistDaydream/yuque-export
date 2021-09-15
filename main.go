package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func LogInit(level, file string) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	le, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(le)

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
	userToken := pflag.String("token", "", "用户 Token,在 https://www.yuque.com/settings/tokens/ 创建")
	userCookie := pflag.String("cookie", "", "用户 Cookie,通过浏览器的 F12 查看")
	pflag.Parse()

	// 初始化日志
	if err := LogInit(*logLevel, *logFile); err != nil {
		logrus.Fatal(errors.Wrap(err, "set log level error"))
	}

	// 获取待导出节点的信息
	discoveredTOCs, err := GetToc(*userToken)
	if err != nil {
		logrus.Error("获取待导出 TOC 信息失败：", err)
	}

	// 逐一导出节点内容
	for _, discoveredTOC := range discoveredTOCs {
		// 获取待导出笔记的 URL
		exportURL, err := GetURLForExportToc(discoveredTOC, *userCookie)
		if err != nil {
			logrus.Error("获取导出 TOC 的 URL 失败：", err)
		}

		logrus.Info("待导出 TOC 的 URL 为：", exportURL)

		// 开始导出笔记
		err = ExportDoc(exportURL, discoveredTOC.Title)
		if err != nil {
			logrus.Error("导出 TOC 失败：", err)
		}
	}

}
