package global

import (
	"gorm.io/gorm"
	"net"
)

const (
	Version  = "v2.0.0"
	LogLevel = 2
	DSN      = "root:123456@tcp(localhost:3306)/acm_game?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	DB              *gorm.DB                    // 数据库
	Listener        net.Listener                // 监听器
	Clients         = make(map[net.Conn]bool)   // 客户端列表
	UserName        = make(map[net.Conn]string) // 用户名列表
	UserNameDisplay = make(map[net.Conn]string) // 用户名列表 (展示用)
	ConnList        = make(map[string]net.Conn) // 通过用户名查找连接

	Score = make(map[net.Conn]int) // 本局得分

	ChangeLog = "" // 更新日志内容

	MsgChan = make(chan Message) // 消息通道

	GameState = 0 // 游戏状态 (0: 等待开始, 1: 游戏中)
	GameMode  = 0 // 游戏模式
	GameTime  = 0 // 游戏结束时间

	Problems       = make([]Problem, 0)              // 题目列表
	ProblemPenalty = make([]int, 0)                  // 题目罚时
	UserSubmission = make(map[net.Conn][]Submission) // 用户提交记录
	WA_ID          = make(map[string]bool)           // 记录已经提交过的 WA 题目 ID

	BeginTime    = int64(0) // 游戏开始时间
	CheckTime    = int64(0) // 最后一次检查时间
	AddScoreTime = int64(0) // 最后一次加分时间

	GameStopCommand = false
)
