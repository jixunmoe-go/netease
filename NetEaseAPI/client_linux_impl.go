package NetEaseAPI

import (
	"encoding/json"
	"fmt"
	"github.com/JixunMoe/netease-api-go/NetEaseAPI/crypto"
	"net/http"
)

type LinuxClientImpl struct{}

func (c *LinuxClientImpl) Request(n *NetEase, result APIResp, method, path string, params interface{}) error {
	object := map[string]interface{}{
		"method": method,
		"params": params,
		"url":    fmt.Sprintf("%s/api%s", n.TunnelURL, path),
	}
	data, err := json.Marshal(object)
	if err != nil {
		return err
	}

	payload := "eparams=" + crypto.LinuxClientEncrypt(data)

	resp, err := n.Request(c, "POST", "/api/linux/forward", payload)
	if err != nil {
		return err
	}

	return result.Deserialize(string(resp))
}

func (c *LinuxClientImpl) ExtendRequest(n *NetEase, req *http.Request) {
	req.Header.Set("Cookie", n.Cookie)
	req.Header.Set("Origin", "orpheus://orpheus")
}
