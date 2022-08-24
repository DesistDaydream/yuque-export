package v1

import (
	"encoding/json"
	"fmt"
	"time"

	core "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v1"
	"github.com/DesistDaydream/yuque-export/pkg/yuquesdk/services/v1/models"
)

type BookService struct {
	Client *core.Client
}

func NewBookService(client *core.Client) *BookService {
	return &BookService{
		Client: client,
	}
}

func (book *BookService) GetDownloadURL(request *models.BookExportRequest, time time.Duration, repoID int) (*models.BookExport, error) {
	var be *models.BookExport
	endpoint := fmt.Sprintf("books/%v/export", repoID)

	// 解析请求体
	reqBodyByte, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reqOpt := core.RequestOption{
		Method: "POST",
		Time:   time,
	}

	_, err = book.Client.Request(endpoint, reqBodyByte, &be, &reqOpt)
	if err != nil {
		return nil, err
	}

	return be, nil
}
