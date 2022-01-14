package yuque

import (
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
)

func NewReposList() *ReposList {
	return &ReposList{}
}

// 从语雀的 API 中获取知识库列表
func (r *ReposList) Get(h *handler.HandlerObject, name string) error {
	endpoint := "/users/" + name + "/repos"

	yc := handler.NewYuqueClient(h.Opts)
	err := yc.Request("GET", endpoint, r)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReposList) Handle(h *handler.HandlerObject) error {
	panic("not implemented") // TODO: Implement
}

func (r *ReposList) DiscoverTocsList(opts *handler.YuqueOpts) string {
	for _, repo := range r.Data {
		if repo.Name == opts.RepoName {
			return fmt.Sprint(repo.ID)
		}
	}
	return ""
}
