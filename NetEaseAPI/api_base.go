package NetEaseAPI

const defaultEndPoint = "https://music.163.com"
const defaultCookie = "osver=%E7%89%88%E6%9C%AC%2010.13.3%EF%BC%88%E7%89%88%E5%8F%B7%2017D47%EF%BC%89; os=osx; appver=1.5.9; channel=netease;"
const defaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"

// NetEase struct is useful.
type NetEase struct {
	BaseURL   string
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
	n.BaseURL = defaultEndPoint
	n.Cookie = defaultCookie
	n.UserAgent = defaultUserAgent
	n.SetRandomIP()
}
