package wechat

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func HttpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func HttpPost(url string, data url.Values, headers map[string]string) ([]byte, error) {
	res, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func HttpPostJson(url string, data interface{}, headers map[string]string) ([]byte, error) {
	// json encode
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// do request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func RandomString() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(time.Now().Format(time.RFC3339Nano))))
}
