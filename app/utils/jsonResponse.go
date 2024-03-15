package utils

import (
	"net/http"
	"xinmiao/app/utils/stateCode"

	"github.com/gin-gonic/gin"
)

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": stateCode.OK,
		"msg":  "OK",
	})
}
