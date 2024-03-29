package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var client = &http.Client{}

func SendJsonRequest(method string, url string, body interface{}, out interface{}) error {
	var buff *bytes.Buffer
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buff = bytes.NewBuffer(payload)
	} else {
		buff = bytes.NewBuffer(nil)
	}
	req, err := http.NewRequest(method, url, buff)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// todo: this only needed when request from bareksa, need a refactor
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// skip the parse body part
	if out == nil {
		return nil
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resBody, out)
	if err != nil {
		return err
	}

	return nil
}
