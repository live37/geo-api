package test

import (
	"fmt"
	"geo-api/service"
	"geo-api/service/impl"
	"testing"
)

func TestGeo(t *testing.T) {
	geoApi := service.GeoService
	geoApi.Init("../")
	data := geoApi.GetIPLocation("8.8.8.8")
	fmt.Printf("%v", data)

}

func TestMaxmind(t *testing.T) {
	geoApi := new(impl.MaxMindService)
	geoApi.Init("../")
	data := geoApi.Geo("8.8.8.8")
	fmt.Printf("%v", data)
}

func TestIp2region(t *testing.T) {
	geoApi := new(impl.Ip2regionService)
	geoApi.Init("../")
	data := geoApi.Geo("8.8.8.8")
	fmt.Printf("%v", data)
}
