package routes

import (
	"web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AssetRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/:source", controllers.GetAssets)
	routerGroup.GET("/:source/:id", controllers.GetAsset)
	routerGroup.POST("/", controllers.CreateUser)
}
