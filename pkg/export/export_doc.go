package export

import (
	"io/ioutil"
	"net/http"
	"os"
)

func ExportDoc(exportURL string, tocName string) error {
	url := exportURL
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	// req.Header.Add("sec-ch-ua", "<sec-ch-ua>")
	// req.Header.Add("sec-ch-ua-mobile", "<sec-ch-ua-mobile>")
	// req.Header.Add("Upgrade-Insecure-Requests", "<Upgrade-Insecure-Requests>")
	// req.Header.Add("User-Agent", "<User-Agent>")
	// req.Header.Add("Accept", "<Accept>")
	// req.Header.Add("Sec-Fetch-Site", "<Sec-Fetch-Site>")
	// req.Header.Add("Sec-Fetch-Mode", "<Sec-Fetch-Mode>")
	// req.Header.Add("Sec-Fetch-User", "<Sec-Fetch-User>")
	// req.Header.Add("Sec-Fetch-Dest", "<Sec-Fetch-Dest>")
	// req.Header.Add("Accept-Language", "<Accept-Language>")

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
