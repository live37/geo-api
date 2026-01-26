package api

type IPVersion string

const (
	IPv4 IPVersion = "ipv4"
	IPv6 IPVersion = "ipv6"
)

type IPLocation struct {
	// 国家
	Country string `json:"country"`
	// 省份
	Province string `json:"province"`
	// 市
	City   string `json:"city"`
	ISP    string `json:"isp"`
	Region string `json:"-"`
}

type GeoApi interface {
	Init(path string)
	Geo(ip string, version IPVersion) *IPLocation
}
