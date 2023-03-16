package openapi

import (
	"manage_address/common/constans"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, constans.OpSuccess())
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, constans.OpFail("data error"))
}
