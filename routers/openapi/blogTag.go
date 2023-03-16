package openapi

import (
	"manage_address/common/constans"
	"manage_address/common/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		argInt, _ := strconv.Atoi(arg)
		state = argInt
		maps["state"] = state
	}

	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	data["lists"] = models.GetTags(pageNum, pageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, constans.OpSuccessData(data))
}
