package v1

import (
	"encoding/json"
	"fmt"
)

type BookService struct {
	Client *ClientV1
	RepoID int
}

func NewBookService(client *ClientV1, repoID int) *BookService {
	return &BookService{
		Client: client,
		RepoID: repoID,
	}
}

func (b *BookService) GetDownloadURL(request *BookExportRequest) (BookExport, error) {
	var be BookExport
	endpoint := fmt.Sprintf("books/%v/export", b.RepoID)

	// 解析请求体
	reqBodyByte, err := json.Marshal(request)
	if err != nil {
		return BookExport{}, err
	}

	_, err = b.Client.Request("POST", endpoint, reqBodyByte, &be)
	if err != nil {
		return BookExport{}, err
	}

	return be, nil
}
