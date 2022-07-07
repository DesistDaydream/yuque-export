package main

import (
	"github.com/DesistDaydream/yuque-export/pkg/export"
	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/logging"
	"github.com/DesistDaydream/yuque-export/pkg/utils/config"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk"
	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

//
func DiscoverTocs(h *handler.HandlerObject, t core.RepoToc) []core.RepoTocData {
	var discoveredTocs []core.RepoTocData
	// 根据用户设定，筛选出需要导出的文档
	logrus.Infof("当前知识库共有 %v 个节点", len(t.Data))
	for _, data := range t.Data {
		if data.Depth == h.Flags.TocDepth {
			discoveredTocs = append(discoveredTocs, data)
		}
	}

	return discoveredTocs
}

// 导出文档集合
func exportSet(h *handler.HandlerObject, tocsList core.RepoToc) {
	// 发现需要导出的文档
	discoveredTocs := DiscoverTocs(h, tocsList)
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
func exportAll(h *handler.HandlerObject, tocsList core.RepoToc) {
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
func getDocDetail(h *handler.HandlerObject, tocsList core.RepoToc) {
	logrus.Infof("需要导出 %v 篇文档", len(tocsList.Data))

	eds := export.GetDocDetail(h, tocsList.Data)

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
	logFlags := &logging.LoggingFlags{}
	logFlags.AddFlags()
	yhFlags := &handler.YuqueHandlerFlags{}
	yhFlags.AddFlag()
	pflag.Parse()

	// 初始化日志
	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	auth := config.NewAuthInfo("lichenhao.yaml")

	// 通过 sdk 实例化语雀客户端
	y := yuquesdk.NewService(auth.Token)

	// 实例化处理器
	h := handler.NewHandlerObject(*yhFlags, y)

	// 获取用户名称
	userInfo, err := y.User.Get("")
	if err != nil {
		logrus.Fatalln(err)
	} else {
		h.UserName = userInfo.Data.Name
	}

	logrus.Println(userInfo.Data.Name)

	// 获取知识库列表
	repos, err := y.Repo.List(userInfo.Data.Name, "", nil)
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Debugf("%v 用户共有 %v 个知识库", userInfo.Data.Name, len(repos.Data))

	// 获取待导出知识库
	for _, repo := range repos.Data {
		if repo.Name == auth.RepoName {
			h.Namespace = repo.Namespace
			logrus.Infof("将要导出【%v】知识库，Namespace 为 %v", auth.RepoName, repo.Namespace)
			break
		}
	}

	// 获取 Toc 列表
	tocs, err := y.Repo.GetToc(h.Namespace)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("【%v】知识库共有 %v 篇笔记", auth.RepoName, len(tocs.Data))

	switch yhFlags.ExportMethod {
	case "set":
		exportSet(h, tocs)
	case "all":
		exportAll(h, tocs)
	case "get":
		getDocDetail(h, tocs)
	default:
		panic("请指定导出方式")
	}
}
