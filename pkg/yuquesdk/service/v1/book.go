package v1

import (
	"encoding/json"
	"fmt"
	"time"

	corev1 "github.com/DesistDaydream/yuque-export/pkg/yuquesdk/core/v1"
	"github.com/sirupsen/logrus"
)

type BookService struct {
	Client *corev1.Client
}

func NewBookService(client *corev1.Client) *BookService {
	return &BookService{
		Client: client,
	}
}

func (book *BookService) GetDownloadURL(request *corev1.BookExportRequest, time time.Duration, repoID int) (corev1.BookExport, error) {
	logrus.Println("lalalal")
	var be corev1.BookExport
	endpoint := fmt.Sprintf("books/%v/export", repoID)

	// 解析请求体
	reqBodyByte, err := json.Marshal(request)
	if err != nil {
		return corev1.BookExport{}, err
	}

	reqOpt := corev1.RequestOption{
		Method: "POST",
		Time:   time,
	}

	_, err = book.Client.Request(endpoint, reqBodyByte, &be, &reqOpt)
	if err != nil {
		return corev1.BookExport{}, err
	}

	return be, nil
}
