package handler

import "github.com/spf13/pflag"

// YuqueUserOpts 通过命令行标志传递的认证选项
type YuqueUserOpts struct {
	UserName string
	RepoName string
	Cookie   string
	Referer  string
	Token    string
	// 待导出知识库的深度。也就是目录层级
	TocDepth int
	IsExport bool
}

// AddFlag 用来为语雀用户数据设置一些值
func (opt *YuqueUserOpts) AddFlag() {
	pflag.StringVar(&opt.UserName, "user-name", "DesistDaydream", "用户名称")
	pflag.StringVar(&opt.RepoName, "repo-name", "学习知识库", "待导出知识库名称")
	pflag.StringVar(&opt.Token, "user-token", "", "用户 Token,在 https://www.yuque.com/settings/tokens/ 创建")
	pflag.StringVar(&opt.Cookie, "user-cookie", "", "用户 Cookie,通过浏览器的 F12 查看")
	pflag.StringVar(&opt.Referer, "referer", "https://www.yuque.com", "用于获取导出笔记的 URL")
	pflag.IntVar(&opt.TocDepth, "toc-depth", 2, "知识库的深度，即从哪一级目录开始导出")
	pflag.BoolVar(&opt.IsExport, "export", false, "是否真实导出笔记，默认不导出，仅查看可以导出的笔记")
}
