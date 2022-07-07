package v1

import (
	"encoding/json"
	"fmt"

	"github.com/DesistDaydream/yuque-export/pkg/utils/config"
)

type BookService struct {
	client *YuqueClient
}

func NewBookService(client *YuqueClient) *BookService {
	return &BookService{
		client: client,
	}
}

func (b *BookService) Get(tocNodeUUID string, namespace int, authInfo config.AuthInfo) (BookExport, error) {
	var e BookExport
	endpoint := fmt.Sprintf("books/%v/export", namespace)
	// 根据节点信息，配置当前待导出节点的请求体信息
	// 解析请求体
	reqBodyByte, err := json.Marshal(BookExportPost{
		Type:  "lakebook",
		Force: 0,
		// Title:        toc.Title,
		// TocNodeUUID:  toc.UUID,
		// TocNodeURL:   toc.URL,
		TocNodeUUID:  tocNodeUUID,
		WithChildren: true,
	})
	if err != nil {
		return BookExport{}, err
	}

	yc := NewYuqueClient(authInfo)

	_, err = yc.Request("POST", endpoint, reqBodyByte, &e)
	if err != nil {
		return BookExport{}, err
	}

	return e, nil
}
