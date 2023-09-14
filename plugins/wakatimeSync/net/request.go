package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handlePathToken(path string, token string) string {
	return path + "?api_key=" + token
}

func Get[R interface{}](path string, params map[string]interface{}, resp R) error {
	path = handlePathToken(path, params["token"].(string))
	delete(params, "token")

	for key, val := range params {
		path += "&" + key + "=" + val.(string)
	}

	var req, err = http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	return client[R](req, resp)
}

func Post[R interface{}](path string, params map[string]interface{}, resp R) error {
	path = handlePathToken(path, params["token"].(string))

	var body, _ = json.Marshal(params)
	var req, err = http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	return client[R](req, resp)
}

func client[R interface{}](req *http.Request, body R) error {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(respBody) == "{\"error\": \"Unauthorized\"}" {
		return fmt.Errorf("Unauthorized")
	}

	if err := json.Unmarshal(respBody, &body); err != nil {
		return err
	}

	return nil

}
