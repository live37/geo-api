package impl

import (
	"fmt"
	"geo-api/service/api"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"strings"
)

type Ip2regionService struct {
	db *xdb.Searcher
}

func (s *Ip2regionService) Init(path string) {
	dbPath := fmt.Sprintf("%sstatic/ip2region.xdb", path)
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}
	s.db = searcher

}

func (s *Ip2regionService) Geo(ip string) *api.IPLocation {
	region, err := s.db.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return nil
	}
	loc := s.parseRegion(region)
	return loc

}

func (s *Ip2regionService) parseRegion(regionStr string) *api.IPLocation {
	// 默认分割符是|，通常格式: 国家|区域|省份|城市|ISP
	parts := strings.Split(regionStr, "|")

	loc := &api.IPLocation{
		Country:  "未知",
		Province: "未知",
		City:     "未知",
		ISP:      "未知",
		Region:   regionStr,
	}

	if len(parts) >= 5 {
		// 处理国家信息
		if parts[0] != "" && parts[0] != "0" {
			loc.Country = parts[0]
		}

		// 处理省份信息
		if parts[2] != "" && parts[2] != "0" {
			loc.Province = parts[2]
		}

		// 处理城市信息
		if parts[3] != "" && parts[3] != "0" {
			loc.City = parts[3]
			// 如果城市包含省份信息，尝试提取纯城市名
			if strings.Contains(loc.City, loc.Province) {
				loc.City = strings.Replace(loc.City, loc.Province, "", 1)
			}
		}

		// 处理ISP信息
		if parts[4] != "" && parts[4] != "0" {
			loc.ISP = parts[4]
		}
	}

	return loc
}
