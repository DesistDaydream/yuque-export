package main

import (
	"github.com/DesistDaydream/yuque-export/pkg/export"
	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/logging"
	"github.com/DesistDaydream/yuque-export/pkg/utils/config"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v2/models"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// 发现需要导出的 Tocs
func DiscoverTocs(h *handler.HandlerObject, t models.RepoToc, slugs []string) []models.RepoTocData {
	var discoveredTocs []models.RepoTocData
	// 根据用户设定，筛选出需要导出的文档
	logrus.Infof("当前知识库共有 %v 个节点", len(t.Data))
	// 若指定了文档的 slugs，则发现指定知识库深度的文档中指定 slugs 的文档
	if len(slugs) > 0 {
		for _, slug := range slugs {
			for _, data := range t.Data {
				if data.Depth == h.Flags.TocDepth && data.Slug == slug {
					discoveredTocs = append(discoveredTocs, data)
				}
			}
		}
	} else { // 若没指定文档的 slugs，则只发现指定知识库深度的文档
		for _, data := range t.Data {
			if data.Depth == h.Flags.TocDepth {
				discoveredTocs = append(discoveredTocs, data)
			}
		}
	}

	return discoveredTocs
}

// 导出文档集合
func exportSet(h *handler.HandlerObject, tocsList models.RepoToc, auth config.AuthInfo) {
	// 发现需要导出的文档
	discoveredTocs := DiscoverTocs(h, tocsList, auth.Slugs)
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
	export.ExportSet(h, discoveredTocs, auth)

	logrus.WithFields(logrus.Fields{
		"总共": len(discoveredTocs),
		"成功": export.SuccessCount,
		"失败": export.FailureCount,
	}).Info("导出完成，统计报告")
}

// 导出知识库中每篇文档
func exportAll(h *handler.HandlerObject, tocsList models.RepoToc) {
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
func getDocDetail(h *handler.HandlerObject, tocsList models.RepoToc) {
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
	authFile := pflag.StringP("file", "f", "DesistDaydream.yaml", "配置文件路径")
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

	auth := config.NewAuthInfo(*authFile)

	// 通过 sdk 实例化语雀客户端
	y := yuquesdk.NewService(auth.Token)
	y1 := yuquesdk.NewServiceV1(auth)

	// 实例化处理器
	h := handler.NewHandlerObject(*yhFlags, y, y1)

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
			h.RepoID = repo.ID
			logrus.Infof("将要导出【%v】知识库，Namespace 为 %v", auth.RepoName, repo.Namespace)
			break
		}
	}

	if h.Namespace == "" {
		logrus.Fatalln("未找到待导出的知识库")
	}

	// 获取 Toc 列表
	tocs, err := y.Repo.GetToc(h.Namespace)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("【%v】知识库共有 %v 篇笔记", auth.RepoName, len(tocs.Data))

	switch yhFlags.ExportMethod {
	case "set":
		exportSet(h, tocs, *auth)
	case "all":
		exportAll(h, tocs)
	case "get":
		getDocDetail(h, tocs)
	default:
		panic("请指定导出方式")
	}
}
