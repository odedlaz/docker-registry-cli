package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Get = "GET"

func prettyPrintJson(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "  ")
	return out.Bytes(), err
}

func Call(registry, username, password, path, method string) ([]byte, error, http.Response) {
	//
	// Get HTTP Client
	client := &http.Client{}
	url := fmt.Sprintf("https://%s/v2/%s", registry, path)

	// Create Request object
	req, err := http.NewRequest(method, url, nil)
	req.SetBasicAuth(username, password)

	// Make API call
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// Get API response body
	body, err := ioutil.ReadAll(resp.Body)

	return body, err, *resp
}
