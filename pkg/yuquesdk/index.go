package yuquesdk

import (
	"github.com/DesistDaydream/yuque-export/pkg/utils/config"
	corev1 "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v1"
	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk/service"
	servicev1 "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/service/v1"
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

type ServiceV1 struct {
	Client     *corev1.Client
	BookExport *servicev1.BookService
}

func NewServiceV1(auth *config.AuthInfo) *ServiceV1 {
	s := new(ServiceV1)
	s.Init(auth)
	return s
}

func (s *ServiceV1) Init(auth *config.AuthInfo) {
	s.Client = corev1.NewClient(auth)
	s.BookExport = servicev1.NewBookService(s.Client)
}
