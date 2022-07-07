package yuquesdk

import (
	"errors"
	"fmt"
)

// DocService encapsulate authenticated token
type DocService struct {
	client *Client
}

// NewDoc create Doc for external use
func NewDoc(client *Client) *DocService {
	return &DocService{
		client: client,
	}
}

// List doc of a repo
func (doc DocService) List(namespace string) (BookDetail, error) {
	var b BookDetail
	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	url := fmt.Sprintf("repos/%s/docs/", namespace)
	_, err := doc.client.RequestObj(url, &b, EmptyRO)
	if err != nil {
		return b, err
	}
	return b, nil
}

// Get detail info of a doc
func (doc DocService) Get(namespace string, slug string, data *DocGet) (DocDetail, error) {
	var b DocDetail
	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	url := fmt.Sprintf("repos/%s/docs/%s", namespace, slug)
	_, err := doc.client.RequestObj(url, &b, &RequestOption{
		Data: StructToMapStr(data),
	})
	if err != nil {
		return b, err
	}
	return b, nil
}

// Create doc
func (doc DocService) Create(namespace string, data *DocCreate) (DocDetail, error) {
	var b DocDetail
	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	if len(data.Format) == 0 {
		data.Format = "markdown"
	}
	url := fmt.Sprintf("repos/%s/docs", namespace)
	_, err := doc.client.RequestObj(url, &b, &RequestOption{
		Method: "POST",
		Data:   StructToMapStr(data),
	})
	if err != nil {
		return b, err
	}
	return b, nil
}

// Update doc
func (doc DocService) Update(namespace string, id string, data *DocCreate) (DocDetail, error) {
	var b DocDetail

	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	if len(id) == 0 {
		return b, errors.New("doc id is required")
	}
	url := fmt.Sprintf("repos/%s/docs/%s", namespace, id)
	_, err := doc.client.RequestObj(url, &b, &RequestOption{
		Method: "PUT",
		Data:   StructToMapStr(data),
	})
	if err != nil {
		return b, err
	}
	return b, nil
}

// Delete doc
func (doc DocService) Delete(namespace string, id string) (DocDetail, error) {
	var b DocDetail
	if len(namespace) == 0 {
		return b, errors.New("repo namespace or id is required")
	}
	if len(id) == 0 {
		return b, errors.New("doc id is required")
	}
	url := fmt.Sprintf("repos/%s/docs/%s", namespace, id)
	_, err := doc.client.RequestObj(url, &b, &RequestOption{
		Method: "DELETE",
	})
	if err != nil {
		return b, err
	}
	return b, nil
}
