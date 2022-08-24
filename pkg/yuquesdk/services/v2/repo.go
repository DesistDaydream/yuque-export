package v2

import (
	"errors"
	"fmt"

	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v2/models"
)

// RepoService encapsulate authenticated token
type RepoService struct {
	client *core.Client
}

// NewRepo create User for external use
func NewRepo(client *core.Client) *RepoService {
	return &RepoService{
		client: client,
	}
}

// List url
func (r RepoService) List(user string, group string, data map[string]string) (*models.UserRepos, error) {
	var (
		url string
		u   *models.UserRepos
	)
	if len(user) == 0 && len(group) == 0 {
		return nil, errors.New("user or group is required")
	}
	if len(user) > 0 {
		url = fmt.Sprintf("users/%s/repos", user)
	} else {
		url = fmt.Sprintf("groups/%s/repos", group)
	}
	_, err := r.client.RequestObj(url, &u, &core.RequestOption{
		Data: data,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Create repo
func (r RepoService) Create(user string, group string, cr *models.CreateRepo) (*models.CreateUserRepo, error) {
	var (
		url string
		u   *models.CreateUserRepo
	)
	if len(user) == 0 && len(group) == 0 {
		return nil, errors.New("user or group is required")
	}
	if len(user) > 0 {
		url = fmt.Sprintf("users/%s/repos", user)
	} else {
		url = fmt.Sprintf("groups/%s/repos", group)
	}
	_, err := r.client.RequestObj(url, &u, &core.RequestOption{
		Method: "POST",
		Data:   StructToMapStr(cr),
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Get repo
func (r RepoService) Get(namespace string, t string) (*models.CreateUserRepo, error) {
	var u *models.CreateUserRepo

	if len(namespace) == 0 && len(t) == 0 {
		return nil, errors.New("namespace or type is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &core.RequestOption{
		Data: map[string]string{"type": t},
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Update repo
func (r RepoService) Update(namespace string, cr *models.UpdateRepo) (*models.CreateUserRepo, error) {
	var u *models.CreateUserRepo

	if len(namespace) == 0 {
		return nil, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &core.RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(cr),
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Delete repo
func (r RepoService) Delete(namespace string) (*models.CreateUserRepo, error) {
	var u *models.CreateUserRepo
	if len(namespace) == 0 {
		return nil, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &core.RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

//GetToc of repo
func (r RepoService) GetToc(namespace string) (*models.RepoToc, error) {
	var u *models.RepoToc
	if len(namespace) == 0 {
		return nil, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s/toc/", namespace)
	_, err := r.client.RequestObj(url, &u, core.EmptyRO)
	if err != nil {
		return nil, err
	}
	return u, nil
}
