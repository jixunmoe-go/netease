package NetEaseAPI

import (
	"encoding/json"
	"fmt"
	"github.com/JixunMoe/netease-api-go/NetEaseAPI/crypto"
	"net/http"
	"net/url"
)

type WEAPIClientImpl struct{}

func (c *WEAPIClientImpl) Request(n *NetEase, result APIResp, method, path string, params interface{}) error {
	data, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to convert input to json string for weapi: %s", err)
	}
	crypted, err := crypto.EncryptWEAPIPayload(data)
	if err != nil {
		return fmt.Errorf("failed to encrypt for weapi: %s", err)
	}

	values := url.Values{}
	values.Add("params", crypted.EncryptedText)
	values.Add("encSecKey", crypted.SecureKey)
	payload := values.Encode()

	gateway := fmt.Sprintf("/weapi%s?csrf_token=7fe1f27659ec97ab196d47b120fe82dc", path)
	resp, err := n.Request(c, method, gateway, payload)
	if err != nil {
		return fmt.Errorf("request to weapi failed: %s", err)
	}

	return result.Deserialize(string(resp))
}

func (c *WEAPIClientImpl) ExtendRequest(n *NetEase, req *http.Request) {
	req.Header.Set("Cookie", n.Cookie+" __csrf=7fe1f27659ec97ab196d47b120fe82dc")
	req.Header.Set("Origin", n.BaseURL)
}
