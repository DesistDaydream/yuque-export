package export

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/DesistDaydream/yuque-export/pkg/utils/converter"
	"github.com/DesistDaydream/yuque-export/pkg/yuque"
)

func ExportDoc(exportURL string, path string, tocName string) error {
	url := exportURL
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s/%s.lakebook", path, tocName)

	os.WriteFile(fileName, body, 0666)

	return nil
}

func ExportMd(dd *yuque.DocDetailData, path string) error {
	// 将文档标题中的 / 替换为 -，防止无法创建文件
	newName := strings.ReplaceAll(dd.Data.Title, "/", "-")
	newName = strings.ReplaceAll(newName, " ", "-")
	newName = strings.ReplaceAll(newName, "(", "-")
	newName = strings.ReplaceAll(newName, ")", "")

	fileName := fmt.Sprintf("%s/%s-%s.md", path, newName, dd.Data.Slug)

	md, err := converter.ConvertHTML2Markdown(dd.Data.BodyHTML)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, []byte(md), 0666)
	if err != nil {
		return fmt.Errorf("写入 %v 文件由于 %v 原因而失败", newName, err)
	}

	return nil
}
