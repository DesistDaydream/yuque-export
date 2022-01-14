package export

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/DesistDaydream/yuque-export/pkg/utils/converter"
)

func ExportDoc(exportURL string, tocName string) error {
	url := exportURL
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	// req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fileName := "./files/" + tocName + ".lakebook"

	os.WriteFile(fileName, body, 0666)

	return nil
}

func ExportMd(data string, name string) error {
	mark, err := converter.ConvertHTML2Markdown(data)
	if err != nil {
		return err
	}

	b := []byte(mark)
	fileName := "./files/" + name + ".md"
	os.WriteFile(fileName, b, 0666)

	return nil
}
