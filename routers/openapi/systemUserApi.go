package openapi

import (
	"log"
	"manage_address/common/constans"
	"manage_address/common/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ManagementLoginDTO struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var params ManagementLoginDTO
	err := c.BindJSON(&params)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, constans.OpFail("request body miss"))
		return
	}

	systemUser := models.SelectOneByAccountAndPassword(params.Account, params.Password)
	if systemUser.ID == 0 {
		log.Println(err)
		c.JSON(http.StatusBadRequest, constans.OpFail("login fail, account or password error"))
		return
	}

	c.JSON(http.StatusOK, constans.OpSuccessData(strings.ToUpper(strings.ReplaceAll(uuid.NewString(), "-", ""))))
}
