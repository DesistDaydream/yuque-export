package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

func HttpHandler(method, url, token string, data interface{}) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
