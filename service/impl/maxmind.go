package impl

import (
	"fmt"
	"geo-api/service/api"
	"github.com/oschwald/maxminddb-golang"
	"log"
	"net"
)

type MaxMindService struct {
	db *maxminddb.Reader
}

func (s *MaxMindService) Init(path string) {
	dbPath := fmt.Sprintf("%sstatic/GeoLite2-City.mmdb", path)
	db, err := maxminddb.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	s.db = db

}

func (s *MaxMindService) Geo(ip string) *api.IPLocation {
	type GeoRecord struct {
		Country struct {
			ISOCode string            `maxminddb:"iso_code"`
			Names   map[string]string `maxminddb:"names"`
		} `maxminddb:"country"`
		City struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"city"`
		Location struct {
			Latitude  float64 `maxminddb:"latitude"`
			Longitude float64 `maxminddb:"longitude"`
			TimeZone  string  `maxminddb:"time_zone"`
		} `maxminddb:"location"`
		// 可以根据需要添加更多字段，如 Continent, Subdivisions 等
	}

	// 要查询的 IP 地址
	ipType := net.ParseIP(ip)
	var record GeoRecord

	// 2. 查询 IP
	err := s.db.Lookup(ipType, &record)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
