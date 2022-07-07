package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
)

// 实例化一个仓库列表
func NewReposList() *ReposList {
	return &ReposList{}
}

// 从语雀的 API 中获取知识库列表
func (r *ReposList) Get(h *handler.HandlerObject, name string) error {
	endpoint := fmt.Sprintf("/users/%s/repos", name)

	yc := handler.NewYuqueClient(h.Flags)
	err := yc.RequestV2("GET", endpoint, nil, r)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReposList) Handle(h *handler.HandlerObject) error {
	panic("not implemented") // TODO: Implement
}

// 发现需要导出的知识库
func (r *ReposList) DiscoverRepos(opts *handler.YuqueHandlerFlags) string {
	for _, repo := range r.Data {
		if repo.Name == opts.RepoName {
			return fmt.Sprint(repo.ID)
		}
	}
	return ""
}
