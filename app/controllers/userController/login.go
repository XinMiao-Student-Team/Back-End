package usercontroller

import (
	"xinmiao/app/apiException"
	"xinmiao/app/services/sessionServices"
	"xinmiao/app/services/userServices"
	"xinmiao/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type autoLoginForm struct {
	Code      string `json:"code" binding:"required"`
	LoginType string `json:"type"`
}
type passwordLoginForm struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	LoginType string `json:"type"`
}

func LoginByPassword(c *gin.Context) {
	// 通过密码登录
	var postForm passwordLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 验证用户
	user, err := userServices.GetUserByStudentIDAndPassword(postForm.Username, postForm.Password)
	if err == gorm.ErrRecordNotFound {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	// 发放session
	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 自定义响应
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"studentID": user.StudentID,
			// 其他字段
			"username":   user.Username,
			"school":     user.School,
			"from":       user.From,
			"grade":      user.Grade,
			"profession": user.Profession,
		},
	})
}

func AuthBySession(c *gin.Context) {
	user, err := sessionServices.UpdateUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"studentID": user.StudentID,
			// 其他字段
			"username":   user.Username,
			"school":     user.School,
			"from":       user.From,
			"grade":      user.Grade,
			"profession": user.Profession,
		},
	})
}
