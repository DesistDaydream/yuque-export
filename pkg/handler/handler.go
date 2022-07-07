package handler

import (
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk"
	"github.com/spf13/pflag"
)

// YuqueHandlerFlags 通过命令行标志传递的认证选项
type YuqueHandlerFlags struct {
	// 导出方式
	ExportMethod string
	// 导出的文件保存路径
	Path string

	// 关于 http.Client 的选项
	// HTTP 请求的超时时间
	Timeout time.Duration

	// 导出每篇笔记的间隔时间，防止并发过大，语雀将会拒绝请求
	ExportDuration int64
	// 并发数
	Concurrency int

	// 专用于导出文档集合的选项
	// 待导出知识库的深度。也就是目录层级
	TocDepth int
	// 是否导出笔记，用来测试
	IsExport bool
}

// AddFlag 用来为语雀用户数据设置一些值
func (opts *YuqueHandlerFlags) AddFlag() {
	pflag.StringVar(&opts.ExportMethod, "method", "set", "导出方式,one of: set|all.set 导出文档集合;all 导出每一篇文档")
	pflag.StringVar(&opts.Path, "paht", "./files", "导出路径")

	pflag.DurationVar(&opts.Timeout, "time-out", time.Second*60, "Timeout on HTTP requests to the Yuque API.unit:second")

	pflag.IntVar(&opts.TocDepth, "toc-depth", 2, "知识库的深度，即从哪一级目录开始导出")
	pflag.BoolVar(&opts.IsExport, "export", false, "是否真实导出笔记，默认不导出，仅查看可以导出的笔记")
	pflag.Int64Var(&opts.ExportDuration, "export-duration", 15, "导出每篇笔记的间隔时间，防止并发过大，语雀将会拒绝请求")
	pflag.IntVar(&opts.Concurrency, "concurrency", 1, "并发数量.")
}

// 用来处理语雀API的数据
type HandlerObject struct {
	// 通过 Token 获取到的用户名称
	UserName string
	// 待导出的知识库。可以是仓库的ID，也可以是以斜线分割的用户名和仓库slug的组合
	Namespace string
	RepoID    int

	// 命令行选项
	Flags YuqueHandlerFlags

	Client   *yuquesdk.Service
	ClientV1 *yuquesdk.ServiceV1
}

// 根据命令行标志实例化一个处理器
func NewHandlerObject(flags YuqueHandlerFlags, client *yuquesdk.Service, clientv1 *yuquesdk.ServiceV1) *HandlerObject {
	return &HandlerObject{
		Flags:    flags,
		Client:   client,
		ClientV1: clientv1,
	}
}
