package impl

import (
	"fmt"
	"geo-api/service/api"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type Ip2regionService struct {
	ipv4Db *xdb.Searcher
	ipv6Db *xdb.Searcher
}

func (s *Ip2regionService) Init(path string) {
	//ip2region, err := service.NewIp2Region(v4Config, v6Config)
	v4Searcher, err := s.initDb(fmt.Sprintf("%sstatic/ip2region_v4.xdb", path), xdb.IPv4)
	if err != nil {
		fmt.Printf("failed to create v4Searcher: %s\n", err.Error())
	}
	v6Searcher, err := s.initDb(fmt.Sprintf("%sstatic/ip2region_v6.xdb", path), xdb.IPv6)
	if err != nil {
		fmt.Printf("failed to create v6Searcher: %s\n", err.Error())
	}
	s.ipv4Db = v4Searcher
	s.ipv6Db = v6Searcher

}
func (s *Ip2regionService) initDb(dbPath string, version *xdb.Version) (*xdb.Searcher, error) {
	// 1、从 dbPath 加载整个 xdb 到内存
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
		return nil, err
	}

	// 2、用全局的 cBuff 创建完全基于内存的查询对象。
	searcher, err := xdb.NewWithBuffer(version, cBuff)
	if err != nil {
		fmt.Printf("failed to create searcher with content: %s\n", err)
		return nil, err
	}
	return searcher, nil
}

func (s *Ip2regionService) Geo(ip string, version api.IPVersion) *api.IPLocation {
	var db *xdb.Searcher
	if version == api.IPv4 {
		db = s.ipv4Db
	} else {
		db = s.ipv6Db
	}
	region, err := db.SearchByStr(ip)
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
		if parts[1] != "" && parts[1] != "0" {
			loc.Province = parts[1]
		}

		// 处理城市信息
		if parts[2] != "" && parts[2] != "0" {
			loc.City = parts[2]
			// 如果城市包含省份信息，尝试提取纯城市名
			if strings.Contains(loc.City, loc.Province) {
				loc.City = strings.Replace(loc.City, loc.Province, "", 1)
			}
		}

		// 处理ISP信息
		if parts[3] != "" && parts[3] != "0" {
			loc.ISP = parts[4]
		}
	}

	return loc
}
