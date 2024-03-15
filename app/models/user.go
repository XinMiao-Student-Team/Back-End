package models

import "time"

type User struct {
	// 因为考虑和管理员公用数据库，所以这里打算用ID来作为账号的唯一标识符
	ID         string // 账号唯一标识符
	StudentID  string
	Username   string    // 用户名
	RealName   string    // 实际姓名
	Password   string    // 数据库
	School     string    // 学校
	From       string    // 生源地
	Type       UserType  // 用户类型
	Grade      int       // 年级
	Profession string    // 专业
	Email      string    // 邮箱
	Telephone  string    // 手机号码
	CreateTime time.Time // 创建时间
}

type UserType int

const (
	Student    UserType = 0 // 学生
	Counsellor UserType = 1 // 辅导员
	Admin      UserType = 2 // 学校后台管理员
	Maintainer UserType = 3 // 运维管理员
)
