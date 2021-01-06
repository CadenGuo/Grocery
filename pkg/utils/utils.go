package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func CreatePostJsonRequest(
	url string,
	body map[string]interface{},
	headers map[string]string,
) *http.Request {
	bodyByte, _ := json.Marshal(body)
	_body := bytes.NewBuffer(bodyByte)
	req, _ := http.NewRequest(http.MethodPost, url, _body)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

func CreatePutJsonRequest(
	url string,
	body map[string]interface{},
	headers map[string]string,
) *http.Request {
	bodyByte, _ := json.Marshal(body)
	_body := bytes.NewBuffer(bodyByte)
	req, _ := http.NewRequest(http.MethodPut, url, _body)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

func CreateGetJsonRequest(
	url string,
	query map[string]string,
	headers map[string]string,
) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req
}

func CallRequestWantJson(request *http.Request) (interface{}, error) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 || resp.StatusCode <= 100 {
		return nil, errors.New(string(ret))
	}
	var jsonBody = make(map[string]interface{})
	err = json.Unmarshal(ret, &jsonBody)
	if err != nil {
		return nil, err
	}
	return jsonBody, nil
}
