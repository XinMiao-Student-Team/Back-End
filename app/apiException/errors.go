package apiException

import "net/http"

type Error struct {
	// Error 的时候一般没有返回data字段，因此省略
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	NotFound              = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	UserNotFind           = NewError(http.StatusInternalServerError, 200503, "该用户不存在")
	NoThatPasswordOrWrong = NewError(http.StatusInternalServerError, 200504, "密码错误")
	StudentIdError        = NewError(http.StatusInternalServerError, 200513, "学号格式不正确，请重新输入")
	ServerError           = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	ParamError            = NewError(http.StatusInternalServerError, 200501, "参数错误")
	NotLogin              = NewError(http.StatusInternalServerError, 200503, "未登录")
	SendVerifyCodeError   = NewError(http.StatusInternalServerError, 200507, "发送验证码失败")
	WrongVerificationCode = NewError(http.StatusInternalServerError, 200509, "验证码错误")
	Unknown               = NewError(http.StatusInternalServerError, 300500, "系统异常，请稍后重试!")
)

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

func (e *Error) Error() string {
	return e.Msg
}

func OtherError(msg string) *Error {
	return NewError(http.StatusForbidden, 100403, msg)
}
