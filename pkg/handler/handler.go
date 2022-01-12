package handler

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// 用来处理语雀API的数据
type HandlerObject struct {
	// 待导出的知识库。可以是仓库的ID，也可以是以斜线分割的用户名和仓库slug的组合
	Namespace int
	Opts      YuqueUserOpts
}

// 根据命令行标志实例化一个处理器
func NewHandlerObject(opts YuqueUserOpts) *HandlerObject {
	return &HandlerObject{
		Opts: opts,
	}
}

// 从语雀的 API 中获取用户数据
func (h *HandlerObject) GetUserData() (*UserData, error) {
	var user UserData
	url := YuqueBaseAPI + YuqueUserAPI
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debug("检查 URL，获取用户数据")

	_, err := HttpHandler("GET", url, h.Opts.Token, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// 从语雀的 API 中获取知识库列表
func (h *HandlerObject) GetReposList() (*ReposList, error) {
	var repos ReposList
	url := YuqueBaseAPI + "/users/" + h.Opts.UserName + YuqueReposAPI
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debug("检查 URL，获取知识库列表")

	_, err := HttpHandler("GET", url, h.Opts.Token, &repos)
	if err != nil {
		return nil, err
	}

	return &repos, err
}

// 从语雀的 API 中获取知识库内的文档列表
func (h *HandlerObject) GetTocsList() (*TocsList, error) {
	var toc TocsList
	url := YuqueBaseAPI + "/repos/" + fmt.Sprint(h.Namespace) + "/toc"
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debug("检查 URL，获取 TOC 数据")

	_, err := HttpHandler("GET", url, h.Opts.Token, &toc)
	if err != nil {
		return nil, err
	}

	return &toc, err
}

// 根据用户设定，筛选出需要导出的文档
func (h *HandlerObject) DiscoveredTocs() ([]TOC, error) {
	var (
		discoveredTOCs []TOC // 已发现节点
	)

	// 获取待导出节点的信息
	tocs, err := h.GetTocsList()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("获取待导出 TOC 信息失败!")
	}

	logrus.Infof("当前知识库共有 %v 个节点", len(tocs.Data))

	for i := 0; i < len(tocs.Data); i++ {
		if tocs.Data[i].Depth == h.Opts.TocDepth {
			discoveredTOCs = append(discoveredTOCs, tocs.Data[i])
		}
	}

	// 输出一些 Debug 信息
	logrus.Infof("已发现 %v 个节点", len(discoveredTOCs))

	for _, discoveredTOC := range discoveredTOCs {
		// logrus.WithFields(logrus.Fields{"toc": discoveredTOC.Title}).Debugf("显示已发现的节点")
		logrus.WithFields(logrus.Fields{
			"title":         discoveredTOC.Title,
			"toc_node_uuid": discoveredTOC.UUID,
			"toc_node_url":  discoveredTOC.URL,
		}).Debug("显示已发现的节点信息")
	}

	return discoveredTOCs, nil
}
