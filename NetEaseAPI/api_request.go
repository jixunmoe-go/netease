package NetEaseAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ClientExtension interface {
	ExtendRequest(n *NetEase, req *http.Request)
}

func (n *NetEase) Request(ext ClientExtension, action, path, data string) ([]byte, error) {
	isPost := action == "POST"
	url := fmt.Sprintf("%s%s", n.BaseURL, path)

	var bodyReader *strings.Reader = nil
	if isPost {
		bodyReader = strings.NewReader(data)
	}

	req, err := http.NewRequest(action, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", n.UserAgent)
	req.Header.Set("X-Real-IP", n.FakeIP)
	req.Header.Set("Referer", n.BaseURL)
	req.Header.Set("Accept-language", "zh-cn")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "orpheus://orpheus")
	req.Header.Set("Cookie", n.Cookie)

	if isPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Content-Length", string(len(data)))
	}

	// Last chance to change anything...!
	if ext != nil {
		ext.ExtendRequest(n, req)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()

	return result, err
}
