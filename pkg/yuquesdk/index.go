package yuquesdk

import (
	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk/service"
)

// Service encapsulate authenticated token
type Service struct {
	Client *core.Client
	User   *service.UserService
	Doc    *service.DocService
	Repo   *service.RepoService
	Group  *service.GroupService
}

// NewService create Client for external use
func NewService(token string) *Service {
	s := new(Service)
	s.Init(token)
	return s
}

func (s *Service) Init(token string) {
	s.Client = core.NewClient(token)
	s.User = service.NewUser(s.Client)
	s.Doc = service.NewDoc(s.Client)
	s.Repo = service.NewRepo(s.Client)
	s.Group = service.NewGroup(s.Client)
}
