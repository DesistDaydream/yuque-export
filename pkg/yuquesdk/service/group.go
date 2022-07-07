package service

import (
	"errors"
	"fmt"

	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
)

// GroupService encapsulate authenticated token
type GroupService struct {
	client *core.Client
}

// NewGroup create Doc for external use
func NewGroup(client *core.Client) *GroupService {
	return &GroupService{
		client: client,
	}
}

// List groups
func (g GroupService) List(login string) (core.Groups, error) {
	var (
		url    string
		groups core.Groups
	)
	if len(login) > 0 {
		url = fmt.Sprintf("users/%s/groups", login)
	} else {
		url = "groups"
	}
	_, err := g.client.RequestObj(url, &groups, core.EmptyRO)
	if err != nil {
		return groups, err
	}
	return groups, nil
}

// Get group
func (g GroupService) Get(login string) (core.GroupDetail, error) {
	var gd core.GroupDetail
	if len(login) == 0 {
		return gd, errors.New("group login or id is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &gd, core.EmptyRO)
	if err != nil {
		return gd, err
	}
	return gd, nil
}

// Create group
func (g GroupService) Create(cg *core.CreateGroup) (core.GroupDetail, error) {
	var gd core.GroupDetail
	if len(cg.Name) == 0 {
		return gd, errors.New("data.name is required")
	}
	if len(cg.Login) == 0 {
		return gd, errors.New("data.login is required")
	}
	_, err := g.client.RequestObj("groups", &gd, &core.RequestOption{
		Method: "POST",
		Data:   StructToMapStr(cg),
	})
	if err != nil {
		return gd, err
	}
	return gd, nil
}

// Update group
func (g GroupService) Update(login string, cg *core.CreateGroup) (core.GroupDetail, error) {
	var groups core.GroupDetail

	if len(login) == 0 {
		return groups, errors.New("group login or id is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &groups, &core.RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(cg),
	})
	if err != nil {
		return groups, err
	}
	return groups, nil
}

// Delete group
func (g GroupService) Delete(login string) (core.GroupDetail, error) {
	var groups core.GroupDetail
	if len(login) == 0 {
		return groups, errors.New("group login or id is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s", login), &groups, &core.RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return groups, err
	}
	return groups, nil
}

// ListUsers of group
func (g GroupService) ListUsers(login string) (core.GroupUsers, error) {
	var gd core.GroupUsers
	if len(login) == 0 {
		return gd, errors.New("group login or id is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users", login), &gd, core.EmptyRO)
	if err != nil {
		return gd, err
	}
	return gd, nil
}

// ListUsers of group
func (g GroupService) AddUser(group string, user string, ga *core.GroupAddUser) (core.GroupUserInfo, error) {
	var gd core.GroupUserInfo

	if len(group) == 0 || len(user) == 0 {
		return gd, errors.New("group and user is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users/%s", group, user), &gd, &core.RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(ga),
	})
	if err != nil {
		return gd, err
	}
	return gd, nil
}

// RemoveUser of group
func (g GroupService) RemoveUser(group string, user string) (core.RemoveUserResponse, error) {
	var gd core.RemoveUserResponse
	if len(group) == 0 || len(user) == 0 {
		return gd, errors.New("group and user is required")
	}
	_, err := g.client.RequestObj(fmt.Sprintf("groups/%s/users/%s", group, user), &gd, &core.RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return gd, err
	}
	return gd, nil
}
