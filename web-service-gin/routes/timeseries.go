package routes

import (
	"log"
	"strconv"
	"web-service-gin/schema"
	"web-service-gin/service"

	"github.com/gin-gonic/gin"
)

func TsRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(c *gin.Context) {
		data := service.GetDummtTSData(10)
		c.JSON(200, data)
	})

	routerGroup.GET("/:n", func(c *gin.Context) {
		n, err := strconv.Atoi(c.Param("n"))
		if err != nil {
			log.Fatal(err)
			return
		}

		data := service.GetDummtTSData(n)
		c.JSON(200, data)
	})

	routerGroup.POST("/", func(c *gin.Context) {
		var tsData []schema.Timeseries
		c.BindJSON(&tsData)
		c.JSON(200, tsData)
	})
}
