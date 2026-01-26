package service

import (
	"geo-api/service/api"
	"geo-api/service/impl"
	"net"
)

type _GeoService struct {
	geoApi api.GeoApi
}

var GeoService = new(_GeoService)

func (s *_GeoService) Init(path string) {
	s.geoApi = new(impl.Ip2regionService)
	s.geoApi.Init(path)

}

func (s *_GeoService) GetIPLocation(ip string) *api.IPLocation {
	// 判断IPv4还是Ipv6
	if ip == "" {
		return &api.IPLocation{
			Country:  "未知",
			Province: "未知",
			City:     "未知",
			ISP:      "未知",
		}
	}
	ipInfo := net.ParseIP(ip)
	version := api.IPv4
	if ipInfo != nil && ipInfo.To4() == nil {
		version = api.IPv6
	}
	location := s.geoApi.Geo(ip, version)
	if location == nil {
		location = &api.IPLocation{
			Country:  "未知",
			Province: "未知",
			City:     "未知",
			ISP:      "未知",
		}
	}
	return location
}
