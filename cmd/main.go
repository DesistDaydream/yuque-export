package main

import (
	"fmt"
	"os"

	"github.com/DesistDaydream/yuque-export/pkg/export"
	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/yuque"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type yuqueExportFlags struct {
	logLevel  string
	logFile   string
	logFormat string
}

func (flags *yuqueExportFlags) AddYuqueExportFlags() {
	pflag.StringVar(&flags.logLevel, "log-level", "info", "The logging level:[debug, info, warn, error, fatal]")
	pflag.StringVar(&flags.logFile, "log-output", "", "the file which log to, default stdout")
	pflag.StringVar(&flags.logFormat, "log-format", "text", "log format,one of: json|text")
}

// LogInit 日志功能初始化，若指定了 log-output 命令行标志，则将日志写入到文件中
func LogInit(level, file, format string) error {
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			// FieldMap:          map[logrus.fieldKey]string{},
			// CallerPrettyfier: func(*runtime.Frame) (string, string) {},
			PrettyPrint: false,
		})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		return fmt.Errorf("请指定正确的日志格式")
	}

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(logLevel)

	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}

	return nil
}

// 导出文档集合
func exportSet(h *handler.HandlerObject, tocsList *yuque.TocsList) {
	// 发现需要导出的文档
	discoveredTocs := tocsList.DiscoverTocs(h)
	// 输出一些 Debug 信息
	logrus.Infof("已发现 %v 个节点", len(discoveredTocs))

	for _, toc := range discoveredTocs {
		logrus.WithFields(logrus.Fields{
			"title":         toc.Title,
			"toc_node_uuid": toc.UUID,
			"toc_node_url":  toc.URL,
		}).Debug("显示已发现 TOC 的信息")
	}

	// 导出多个文档集合
	export.ExportSet(h, discoveredTocs)

	logrus.WithFields(logrus.Fields{
		"总共": len(discoveredTocs),
		"成功": export.SuccessCount,
		"失败": export.FailureCount,
	}).Info("导出完成，统计报告")
}

// 导出知识库中每篇文档
func exportAll(h *handler.HandlerObject, tocsList *yuque.TocsList) {
	// 获取 Docs 列表
	// Docs 列表需要分页，暂时还不知道怎么处理，先通过 Tocs 列表获取 Slug
	// docsList := yuque.NewDocsList()
	// err := docsList.Get(h)
	// if err != nil {
	// 	panic(err)
	// }
	// docsList.Handle(h)
	logrus.Infof("需要导出 %v 篇文档", len(tocsList.Data))

	// 导出知识库中每篇文档
	export.ExportAll(h, tocsList.Data)

	logrus.WithFields(logrus.Fields{
		"总共": len(tocsList.Data),
		"成功": export.SuccessCount,
		"失败": export.FailureCount,
	}).Info("导出完成，统计报告")
}

// 获取文档详情
func getDocDetail(h *handler.HandlerObject, tocsList *yuque.TocsList) {
	logrus.Infof("需要导出 %v 篇文档", len(tocsList.Data))

	// 导出知识库中每篇文档
	eds := export.GetDocDetail(h, tocsList.Data)

	fmt.Println(eds.ExceptionDocs)

	// 输出异常笔记信息
	for _, d := range eds.ExceptionDocs {
		logrus.WithFields(logrus.Fields{
			"title": d.Title,
			"slug":  d.Slug,
		}).Infof("私密笔记信息!")
	}

	logrus.WithFields(logrus.Fields{
		"总共": len(eds.ExceptionDocs),
	}).Info("统计报告")
}

func main() {
	// 设置命令行标志
	yeFlags := &yuqueExportFlags{}
	yeFlags.AddYuqueExportFlags()
	yhFlags := &handler.YuqueHandlerFlags{}
	yhFlags.AddFlag()
	pflag.Parse()

	// 初始化日志
	if err := LogInit(yeFlags.logLevel, yeFlags.logFile, yeFlags.logFormat); err != nil {
		logrus.Fatal(errors.Wrap(err, "set log level error"))
	}

	h := handler.NewHandlerObject(*yhFlags)

	// 获取用户信息
	userData := yuque.NewUserData()
	err := userData.Get(h, "")
	if err != nil {
		panic(err)
	}

	// 获取用户名称
	h.UserName = userData.GetUserName()

	// 获取知识库列表
	reposList := yuque.NewReposList()
	err = reposList.Get(h, h.UserName)
	if err != nil {
		panic(err)
	}

	// 获取待导出知识库
	h.Namespace = reposList.DiscoverRepos(yhFlags)

	// 获取 Toc 列表
	tocsList := yuque.NewTocsList()
	err = tocsList.Get(h, "")
	if err != nil {
		panic(err)
	}

	switch yhFlags.ExportMethod {
	case "set":
		exportSet(h, tocsList)
	case "all":
		exportAll(h, tocsList)
	case "get":
		getDocDetail(h, tocsList)
	default:
		panic("请指定导出方式")
	}

}
