package global

import "net"

type Message struct {
	Text  string   // 内容
	Conn  net.Conn // 发送给谁
	IsAll bool     // 是否群发
}

type Problem struct {
	Url        string `gorm:"type:varchar(100);null;"`
	Difficulty int    `gorm:"type:int;null;"`
	IsUsed     bool   `gorm:"type:tinyint(1);null;"`
}

type User struct {
	UserName string `gorm:"type:varchar(100);null;"`
	Score    int    `gorm:"type:int;null;"`
	Point    int    `gorm:"type:int;null;"`
}

type Text struct {
	K string `gorm:"type:varchar(255);null;"`
	V string `gorm:"type:varchar(255);null;"`
}

type Submission struct {
	SubmissionID string // 提交ID
	Url          string
	Status       string // 提交状态
}
