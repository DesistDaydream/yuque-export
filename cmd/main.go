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
func set(h *handler.HandlerObject) {
	// 获取 Toc 列表
	tocsList := yuque.NewTocsList()
	err := tocsList.Get(h)
	if err != nil {
		panic(err)
	}

	// 处理 Toc 列表，这里暂时是有一个逻辑，就是发现需要导出的文档
	err = tocsList.Handle(h)
	if err != nil {
		panic(err)
	}

	// 导出多个文档集合
	export.RunSet(h)
}

// 导出单篇文档
func one(h *handler.HandlerObject) {
	// 获取 Docs 列表
	// Docs 列表需要分页，暂时还不知道怎么处理，先通过 Tocs 列表获取 Slug
	// docsList := yuque.NewDocsList()
	// err := docsList.Get(h)
	// if err != nil {
	// 	panic(err)
	// }
	// docsList.Handle(h)

	// 获取 Toc 列表
	tocsList := yuque.NewTocsList()
	err := tocsList.Get(h)
	if err != nil {
		panic(err)
	}
	tocsList.GetTocsSlug(h)
	logrus.Debug("DocSlug 列表:", h.DocsSlug)

	// 导出多个单篇文档
	export.RunOne(h)
}

func main() {
	// 设置命令行标志
	flags := &yuqueExportFlags{}
	flags.AddYuqueExportFlags()
	opts := &handler.YuqueUserOpts{}
	opts.AddFlag()
	pflag.Parse()

	// 初始化日志
	if err := LogInit(flags.logLevel, flags.logFile, flags.logFormat); err != nil {
		logrus.Fatal(errors.Wrap(err, "set log level error"))
	}

	h := handler.NewHandlerObject(*opts)

	// 获取用户信息
	userData := yuque.NewUserData()
	err := userData.Get(h)
	if err != nil {
		panic(err)
	}

	// 处理用户数据
	err = userData.Handle(h)
	if err != nil {
		panic(err)
	}

	// 获取知识库列表
	reposList := yuque.NewReposList()
	err = reposList.Get(h)
	if err != nil {
		panic(err)
	}

	// 获取需要导出的知识库 ID
	for _, repo := range reposList.Data {
		if repo.Name == opts.RepoName {
			h.Namespace = repo.ID
		}
	}

	switch opts.ExportMethod {
	case "set":
		set(h)
	case "one":
		one(h)
	default:
		panic("请指定导出方式")
	}

}
