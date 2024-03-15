package usercontroller

import (
	"fmt"
	"xinmiao/app/apiException"
	"xinmiao/app/services/sessionServices"
	"xinmiao/app/services/userServices"
	"xinmiao/app/utils"

	"github.com/gin-gonic/gin"
)

type ResetPassForm struct {
	ID      string `json:"ID"`
	OldPass string `json:"oldpass"`
	NewPass string `json:"newpass"`
}

func ResetPass(c *gin.Context) {
	// 获取表单数据
	var postForm ResetPassForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	// 验证用户
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.ID != postForm.ID {
		_ = c.AbortWithError(200, apiException.StudentIdError)
		return
	}
	status, err := userServices.GetCode(fmt.Sprintf("%s-%s", postForm.ID, apiException.ResetPassword))
	if err != nil || status != "Accepted" {
		_ = c.AbortWithError(200, apiException.WrongVerificationCode)
	}
	userServices.DelCode(fmt.Sprintf("%s-%s", postForm.ID, apiException.ResetPassword)) // 过期验证码

	utils.JsonSuccessResponse(c, nil)
}
