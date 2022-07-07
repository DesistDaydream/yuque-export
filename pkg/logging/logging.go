package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// 日志相关命令行标志
type LoggingFlags struct {
	LogLevel  string
	LogOutput string
	LogFormat string
}

// 添加命令行标志
func (flags *LoggingFlags) AddFlags() {
	pflag.StringVar(&flags.LogLevel, "log-level", "info", "日志级别:[debug, info, warn, error, fatal]")
	pflag.StringVar(&flags.LogOutput, "log-output", "", "日志输出位置，不填默认标准输出 stdout")
	pflag.StringVar(&flags.LogFormat, "log-format", "text", "日志输出格式: [text, json]")
}

// LogInit 日志功能初始化，若指定了 log-output 命令行标志，则将日志写入到文件中
func LogInit(level, file, format string) error {
	switch format {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
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
