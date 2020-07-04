package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JixunMoe/netease-api-go/pkg/crypto"
	"github.com/JixunMoe/netease-api-go/pkg/ip"
	"github.com/JixunMoe/netease-api-go/pkg/types"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// BeiJing China Telecom IP (ISP)
// 106.37.0.0 - 106.39.255.255
var fakeIpBegin = ip.Make(106, 37, 0, 0)
var fakeIpEnd = ip.Make(106, 39, 255, 255)

const defaultCookie = "osver=%E7%89%88%E6%9C%AC%2010.13.3%EF%BC%88%E7%89%88%E5%8F%B7%2017D47%EF%BC%89; os=osx; appver=1.5.9; channel=netease;"
const defaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"

// NetEase struct is useful.
type NetEase struct {
	Cookie    string
	FakeIP    string
	UserAgent string
}

// New creates a new NetEase instance.
func New() NetEase {
	n := NetEase{}
	n.init()
	return n
}

func (n *NetEase) init() {
	n.Cookie = defaultCookie
	n.UserAgent = defaultUserAgent
	n.SetRandomIP()
}

func (n *NetEase) request(action, path, data string) (string, error) {
	isPost := action == "POST"
	url := fmt.Sprintf("http://music.163.com%s", path)

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

	_ = os.Setenv("HTTP_PROXY", "http://127.0.0.1:8888")
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

func (n *NetEase) Song(ids ...uint) (*types.SongResp, error) {
	if len(ids) < 1 {
		return nil, errors.New("NetEase::Song: no id supplied")
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[{id:%d}", ids[0]))
	for i := 1; i < len(ids); i++ {
		sb.WriteString(fmt.Sprintf(",{id:%d}", ids[i]))
	}
	sb.WriteByte(']')

	param, err := encodeParam(
		"POST",
		"http://music.163.com/api/v3/song/detail",
		map[string]interface{}{
			"c": sb.String(),
		},
	)

	if err != nil {
		return nil, err
	}

	var result types.SongResp
	resp, err := n.linuxRequest(param)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(resp), &result)
	return &result, err
}

func (n *NetEase) Playlist(id uint) (*types.PlayListResp, error) {
	param, err := encodeParam(
		"POST",
		"http://music.163.com/api/v3/playlist/detail",
		map[string]uint{
			"id": id,
			"n":  1000,
		},
	)

	if err != nil {
		return nil, err
	}

	var result types.PlayListResp
	resp, err := n.linuxRequest(param)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(resp), &result)
	return &result, err
}

func (n *NetEase) SetRandomIP() {
	n.FakeIP = ip.Format(ip.GetRandomIP(fakeIpBegin, fakeIpEnd))
}
