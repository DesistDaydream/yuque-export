package service

import (
	"errors"
	"fmt"

	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v2"
)

// DocService encapsulate authenticated token
type DocService struct {
	client *core.Client
}

// NewDoc create Doc for external use
func NewDoc(client *core.Client) *DocService {
	return &DocService{
		client: client,
	}
}

// List doc of a repo
func (doc DocService) List(namespace string) (core.BookDetail, error) {
	var b core.BookDetail
	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	url := fmt.Sprintf("repos/%s/docs/", namespace)
	_, err := doc.client.RequestObj(url, &b, core.EmptyRO)
	if err != nil {
		return b, err
	}
	return b, nil
}

// Get detail info of a doc
func (doc DocService) Get(namespace string, slug string, data *core.DocGet) (core.DocDetail, error) {
	var b core.DocDetail
	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	url := fmt.Sprintf("repos/%s/docs/%s", namespace, slug)
	_, err := doc.client.RequestObj(url, &b, &core.RequestOption{
		Data: StructToMapStr(data),
	})
	if err != nil {
		return b, err
	}
	return b, nil
}

// Create doc
func (doc DocService) Create(namespace string, data *core.DocCreate) (core.DocDetail, error) {
	var b core.DocDetail
	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	if len(data.Format) == 0 {
		data.Format = "markdown"
	}
	url := fmt.Sprintf("repos/%s/docs", namespace)
	_, err := doc.client.RequestObj(url, &b, &core.RequestOption{
		Method: "POST",
		Data:   StructToMapStr(data),
	})
	if err != nil {
		return b, err
	}
	return b, nil
}

// Update doc
func (doc DocService) Update(namespace string, id string, data *core.DocCreate) (core.DocDetail, error) {
	var b core.DocDetail

	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	if len(id) == 0 {
		return b, errors.New("doc id is required")
	}
	url := fmt.Sprintf("repos/%s/docs/%s", namespace, id)
	_, err := doc.client.RequestObj(url, &b, &core.RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(data),
	})
	if err != nil {
		return b, err
	}
	return b, nil
}

// Delete doc
func (doc DocService) Delete(namespace string, id string) (core.DocDetail, error) {
	var b core.DocDetail
	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	if len(id) == 0 {
		return b, errors.New("doc id is required")
	}
	url := fmt.Sprintf("repos/%s/docs/%s", namespace, id)
	_, err := doc.client.RequestObj(url, &b, &core.RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return b, err
	}
	return b, nil
}
