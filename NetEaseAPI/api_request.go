package NetEaseAPI

import (
	"encoding/json"
	"fmt"
	"github.com/JixunMoe/netease-api-go/NetEaseAPI/crypto"
	"io/ioutil"
	"net/http"
	"strings"
)

func (n *NetEase) request(action, path, data string) (string, error) {
	isPost := action == "POST"
	url := fmt.Sprintf("%s%s", n.BaseURL, path)

	var bodyReader *strings.Reader = nil
	if isPost {
		bodyReader = strings.NewReader(data)
	}

	req, err := http.NewRequest(action, url, bodyReader)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", n.Cookie)
	req.Header.Set("User-Agent", n.UserAgent)
	req.Header.Set("X-Real-IP", n.FakeIP)
	req.Header.Set("Referer", "https://music.163.com/")

	if isPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Content-Length", string(len(data)))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()

	return string(result), err
}

func (n *NetEase) linuxRequest(param string) (string, error) {
	return n.request("POST", "/api/linux/forward", param)
}

func encodeParam(method, url string, params interface{}) (string, error) {
	object := map[string]interface{}{
		"method": method,
		"params": params,
		"url":    url,
	}
	data, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	return "eparams=" + crypto.EncryptToString(data), nil
}
