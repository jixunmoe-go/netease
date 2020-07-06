package NetEaseAPI

import "github.com/jixunmoe-go/netease/util/ip"

// BeiJing China Telecom IP (ISP)
// 106.37.0.0 - 106.39.255.255
var fakeIpBegin = ip.Make(106, 37, 0, 0)
var fakeIpEnd = ip.Make(106, 39, 255, 255)

func (n *NetEase) SetRandomIP() {
	n.FakeIP = ip.Format(ip.GetRandomIP(fakeIpBegin, fakeIpEnd))
}
