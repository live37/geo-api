package service

import (
	"geo-api/service/api"
	"geo-api/service/impl"
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
	location := s.geoApi.Geo(ip)
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
