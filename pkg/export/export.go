package export

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/DesistDaydream/yuque-export/pkg/utils/converter"
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

func ExportMd(data string, path string, name string) error {
	// 将文档标题中的 / 替换为 -，防止无法创建文件
	newName := strings.ReplaceAll(name, "/", "-")
	newName = strings.ReplaceAll(newName, " ", "-")
	newName = strings.ReplaceAll(newName, "(", "-")
	newName = strings.ReplaceAll(newName, ")", "")

	mark, err := converter.ConvertHTML2Markdown(data)
	if err != nil {
		return err
	}

	b := []byte(mark)
	fileName := fmt.Sprintf("%s/%s.md", path, newName)
	err = os.WriteFile(fileName, b, 0666)
	if err != nil {
		return fmt.Errorf("写入 %v 文件由于 %v 原因而失败", name, err)
	}

	return nil
}
