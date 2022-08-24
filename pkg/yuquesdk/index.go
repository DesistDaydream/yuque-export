package yuquesdk

import (
	"github.com/DesistDaydream/yuque-export/pkg/utils/config"
	corev1 "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v1"
	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
	servicev1 "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v1"
	services "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v2"
)

// Service encapsulate authenticated token
type Service struct {
	Client *core.Client
	User   *services.UserService
	Doc    *services.DocService
	Repo   *services.RepoService
	Group  *services.GroupService
}

// NewService create Client for external use
func NewService(token string) *Service {
	s := new(Service)
	s.Init(token)
	return s
}

func (s *Service) Init(token string) {
	s.Client = core.NewClient(token)
	s.User = services.NewUser(s.Client)
	s.Doc = services.NewDoc(s.Client)
	s.Repo = services.NewRepo(s.Client)
	s.Group = services.NewGroup(s.Client)
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
