package service

import (
	"errors"
	"fmt"

	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
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
func (r RepoService) List(user string, group string, data map[string]string) (core.UserRepos, error) {
	var (
		url string
		u   core.UserRepos
	)
	if len(user) == 0 && len(group) == 0 {
		return u, errors.New("user or group is required")
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
		return u, err
	}
	return u, nil
}

// Create repo
func (r RepoService) Create(user string, group string, cr *core.CreateRepo) (core.CreateUserRepo, error) {
	var (
		url string
		u   core.CreateUserRepo
	)
	if len(user) == 0 && len(group) == 0 {
		return u, errors.New("user or group is required")
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
		return u, err
	}
	return u, nil
}

// Get repo
func (r RepoService) Get(namespace string, t string) (core.CreateUserRepo, error) {
	var u core.CreateUserRepo

	if len(namespace) == 0 && len(t) == 0 {
		return u, errors.New("namespace or type is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &core.RequestOption{
		Data: map[string]string{"type": t},
	})
	if err != nil {
		return u, err
	}
	return u, nil
}

//Update repo
func (r RepoService) Update(namespace string, cr *core.UpdateRepo) (core.CreateUserRepo, error) {
	var u core.CreateUserRepo

	if len(namespace) == 0 {
		return u, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &core.RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(cr),
	})
	if err != nil {
		return u, err
	}
	return u, nil
}

//Delete repo
func (r RepoService) Delete(namespace string) (core.CreateUserRepo, error) {
	var u core.CreateUserRepo
	if len(namespace) == 0 {
		return u, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s", namespace)
	_, err := r.client.RequestObj(url, &u, &core.RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return u, err
	}
	return u, nil
}

//GetToc of repo
func (r RepoService) GetToc(namespace string) (core.RepoToc, error) {
	var u core.RepoToc
	if len(namespace) == 0 {
		return u, errors.New("namespace is required")
	}
	url := fmt.Sprintf("repos/%s/toc/", namespace)
	_, err := r.client.RequestObj(url, &u, core.EmptyRO)
	if err != nil {
		return u, err
	}
	return u, nil
}
