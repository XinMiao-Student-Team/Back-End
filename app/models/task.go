package models

import "time"

type Task struct {
	ID               string    // 任务标识
	Name             string    // 任务标题
	Description      string    // 任务描述
	StartTime        time.Time // 活动开始时间
	EndTime          time.Time // 活动结束时间
	SignUpStart      time.Time // 开始报名时间
	SignUpEnd        time.Time // 终止报名时间
	Distination      Location  // 位置
	ParticipantLimit int       // 参与人数限制
	Force            bool      // 是否必须参加
	Type             bool      // 是否全员参加 -> true 表示全员参加
	MemberList       []string  // 参与人员名单，为空表示全员参加
}

type Location struct {
	Latitude  float64
	Longitude float64
}
