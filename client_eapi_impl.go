package NetEaseAPI

import (
	"encoding/json"
	"fmt"
	"github.com/jixunmoe-go/netease/crypto"
	"github.com/jixunmoe-go/netease/util"
	"net/http"
)

type EAPIClientImpl struct{}

func (c *EAPIClientImpl) Request(n *NetEase, result APIResp, method, path string, params interface{}) error {
	// force reference
	_ = method

	object := params.(map[string]interface{})

	// Cookie as header
	object["header"] = util.ParseCookie(n.Cookie)

	data, err := json.Marshal(params)

	if err != nil {
		return err
	}

	internalPath := fmt.Sprintf("/api%s", path)
	payload := "params=" + crypto.EncryptEAPIRequestPayload(internalPath, data)

	gateway := fmt.Sprintf("/eapi%s", path)
	resp, err := n.Request(c, "POST", gateway, payload)
	if err != nil {
		return err
	}

	// Let's try and decrypt it
	if len(resp) > 0 && resp[0] != '{' {
		resp = crypto.DecryptEAPIResponse(resp)
	}

	return result.Deserialize(string(resp))
}

func (c *EAPIClientImpl) ExtendRequest(n *NetEase, req *http.Request) {
	// force reference
	_ = n

	req.Header.Set("Origin", "orpheus://orpheus")
}
