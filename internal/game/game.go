package game

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/internal/handler"
	"ACM_GAME_V2/internal/repo"
	"ACM_GAME_V2/log/xlog"
	"fmt"
	"time"
)

func Game() {
	// 配置题目
	GetProblem()
	// 初始化
	global.GameState = 1
	global.BeginTime = time.Now().Unix()
	global.CheckTime = time.Now().Unix()
	// 开始游戏
	xlog.Info("游戏开始")
	handler.SendBroadcast(global.SEND_XITONG + " 游戏开始!")
	// 显示题目
	ListProblemsAll()
	// 进入循环
	for {
		// 检查提交
		if time.Now().Unix()-global.CheckTime >= global.CHECK_CD {
			handler.Check()
		}
		//检查加活跃分
		if time.Now().Unix()-global.AddScoreTime >= global.ADD_SCORE_CD {
			for conn := range global.Clients {
				handler.SendMessage(conn, global.YELLOW+" + 2 贡献分 (参加游戏得分)"+global.RESET)
				global.Score[conn] += 2
				//repo.AddScore(global.UserName[conn], 2)
			}
			global.AddScoreTime = time.Now().Unix()
		}
		// 检查结束游戏
		if int(time.Now().Unix()-global.BeginTime) >= global.GameTime {
			break
		}
		if global.GameStopCommand {
			global.GameStopCommand = false
			break
		}
	}
	// 游戏结束
	AC := 0
	penaly := 0
	for i := 0; i < len(global.Problems); i++ {
		if global.ProblemPenalty[i] >= 0 {
			AC++
			penaly += global.ProblemPenalty[i]
		}
	}
	handler.SendBroadcast(global.SEND_XITONG + fmt.Sprintf(" 游戏结束! 共解决题目: %d, 总罚时: %d", AC, penaly))
	for conn := range global.Clients {
		handler.ListScoresCommand(conn)
	}
	for conn := range global.Score {
		handler.SendMessage(conn, global.YELLOW+fmt.Sprintf(" + %d 活跃等级分", global.Score[conn])+global.RESET)
		repo.AddScore(global.UserName[conn], global.Score[conn])
		global.Score[conn] = 0
	}
	time.Sleep(3 * time.Second)
}
