package handler

import (
	"time"

	"github.com/spf13/pflag"
)

// YuqueOpts 通过命令行标志传递的认证选项
type YuqueOpts struct {
	UserName string
	RepoName string
	Cookie   string
	Referer  string
	Token    string
	// 待导出知识库的深度。也就是目录层级
	TocDepth int
	IsExport bool
	// 导出方式
	ExportMethod string
	// 导出每篇笔记的间隔时间，防止并发过大，语雀将会拒绝请求，主要用在导出集合上
	ExportDuration int64
	// 关于 http.Client 的选项,HTTP 请求的超时时间
	Timeout time.Duration
}

// AddFlag 用来为语雀用户数据设置一些值
func (opts *YuqueOpts) AddFlag() {
	pflag.StringVar(&opts.UserName, "user-name", "DesistDaydream", "用户名称")
	pflag.StringVar(&opts.RepoName, "repo-name", "学习知识库", "待导出知识库名称")
	pflag.StringVar(&opts.Token, "user-token", "", "用户 Token,在 https://www.yuque.com/settings/tokens/ 创建")
	pflag.StringVar(&opts.Cookie, "user-cookie", "", "用户 Cookie,通过浏览器的 F12 查看")
	pflag.StringVar(&opts.Referer, "referer", "https://www.yuque.com", "用于获取导出笔记的 URL")
	pflag.Int64Var(&opts.ExportDuration, "export-duration", 15, "导出每篇笔记的间隔时间，防止并发过大，语雀将会拒绝请求")
	pflag.IntVar(&opts.TocDepth, "toc-depth", 2, "知识库的深度，即从哪一级目录开始导出")
	pflag.BoolVar(&opts.IsExport, "export", false, "是否真实导出笔记，默认不导出，仅查看可以导出的笔记")
	pflag.StringVar(&opts.ExportMethod, "method", "set", "导出方式,one of: set|all.set 导出文档集合;all 导出每一篇文档")
	pflag.DurationVar(&opts.Timeout, "time-out", time.Second*60, "Timeout on HTTP requests to the Yuque API.unit:second")
}
