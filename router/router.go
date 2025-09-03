package router

import (
	"geo-api/service"
	"github.com/gin-gonic/gin"
)

func StartHttp(addr string) {
	router := gin.Default()
	router.GET("/geo", func(c *gin.Context) {
		// 获取ip param或者json
		ip := c.Query("ip")
		res := service.GeoService.GetIPLocation(ip)
		c.JSON(200, res)
	})
	_ = router.Run(addr)
}
