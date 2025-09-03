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
	return s.geoApi.Geo(ip)
}
