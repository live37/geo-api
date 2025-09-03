package main

import (
	"geo-api/router"
	"geo-api/service"
)

func main() {

	service.GeoService.Init("./")
	router.StartHttp(":8080")
}
