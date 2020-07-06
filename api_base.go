package netease

const defaultEndpoint = "https://music.163.com"
const defaultTunnelEndpoint = "http://music.163.com"
const defaultCookie = "osver=Mac OS X 10.13.3 (17D47); os=osx; appver=1.5.9; channel=netease;"
const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko)"

type APIResp interface {
	Deserialize(resp string) error
}

type APIClient interface {
	Request(n *NetEase, result APIResp, method, path string, params interface{}) error
}

// NetEase struct is useful.
type NetEase struct {
	BaseURL   string
	TunnelURL string
	Cookie    string
	FakeIP    string
	UserAgent string
	Client    APIClient
}

// New creates a new NetEase instance.
func New() NetEase {
	n := NetEase{}
	n.init()
	return n
}

func (n *NetEase) init() {
	n.BaseURL = defaultEndpoint
	n.TunnelURL = defaultTunnelEndpoint
	n.Cookie = defaultCookie
	n.UserAgent = defaultUserAgent
	n.UseLinuxClient()
	n.SetRandomIP()
}

func (n *NetEase) UseEAPIClient() {
	n.Client = &EAPIClientImpl{}
}

func (n *NetEase) UseLinuxClient() {
	n.Client = &LinuxClientImpl{}
}

func (n *NetEase) UseWEAPIClient() {
	n.Client = &WEAPIClientImpl{}
}
