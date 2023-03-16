package routers

import (
	"manage_address/pkg"
	"manage_address/routers/openapi"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(pkg.RunMode)
	api := r.Group("/api")

	{
		api.POST("/manage/login", openapi.Login)
	}

	return r
}
