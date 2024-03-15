package usercontroller

import (
	"xinmiao/app/apiException"
	"xinmiao/app/services/sessionServices"
	"xinmiao/app/services/userServices"
	"xinmiao/app/utils"

	"github.com/gin-gonic/gin"
)

type sendcodeForm struct {
	MailTo   string `json:"mailto"`
	Category string `json:"category"` // 业务场景
}
type checkcodeForm struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Category string `json:"category"`
}

func SendMailVerifyCode(c *gin.Context) {
	var postForm sendcodeForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	err = userServices.SendCodeByMail(user.ID, postForm.MailTo, postForm.Category)
	if err != nil {
		_ = c.AbortWithError(200, apiException.SendVerifyCodeError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func CheckVerifyCode(c *gin.Context) {
	var postForm checkcodeForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	// 验证身份
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.ID != postForm.ID {
		_ = c.AbortWithError(200, apiException.StudentIdError)
		return
	}
	ok, err := userServices.CheckCode(postForm.ID, postForm.Category, postForm.Code)
	if err != nil || !ok {
		_ = c.AbortWithError(200, apiException.WrongVerificationCode)
	}

	utils.JsonSuccessResponse(c, nil)
}
