package handler

import (
	"ACM_GAME_V2/common/ais"
	"ACM_GAME_V2/global"
	"fmt"
	"net"
	"time"
)

const HELP_MESSAGE = `
---[普通指令]---------------------------

/help - 显示指令列表

/update - 查看最近更新日志

---[游戏指令]---------------------------

/check - 刷新题目状态

/list - 显示题目列表

/score - 显示贡献分列表

----------------------------------------
`

func HandleCommand(conn net.Conn, message string) {
	if message == "/check" {
		if global.GameState == 1 {
			if time.Now().Unix()-global.CheckTime >= global.CHECK_CD_MIN {
				Check()
			} else {
				SendMessage(conn, global.SEND_XITONG+" 检查频率过快，请稍后再试！")
			}
		} else {
			SendMessage(conn, global.DARKWHITE+" 游戏未开始！"+global.RESET)
		}
	} else if message == "/stop" {
		global.GameStopCommand = true
	} else if message == "/list" {
		if global.GameState == 1 {
			ListProblemsCommand(conn)
		} else {
			SendMessage(conn, global.DARKWHITE+" 游戏未开始！"+global.RESET)
		}
	} else if message == "/update" {
		SendMessage(conn, global.ChangeLog)
	} else if message == "/help" {
		SendMessage(conn, HELP_MESSAGE)
	} else if message == "/score" {
		if global.GameState == 1 {
			ListScoresCommand(conn)
		} else {
			SendMessage(conn, global.DARKWHITE+" 游戏未开始！"+global.RESET)
		}
	}
}

func ListProblemsCommand(conn net.Conn) {
	SendMessage(conn, "\n[题目列表]:")
	for i := 0; i < len(global.Problems); i++ {
		if global.ProblemPenalty[i] == -1 {
			SendMessage(conn, fmt.Sprintf("%d. %s (难度: %d)", i+1, global.Problems[i].Url, global.Problems[i].Difficulty))
			SendMessage(conn, ais.CommonAI("你是一个泼辣的大小姐，语气傲娇", fmt.Sprintf("%d. %s (难度: %d)", i+1, global.Problems[i].Url, global.Problems[i].Difficulty)))
		} else {
			SendMessage(conn, global.DARKWHITE+fmt.Sprintf("%d. %s (难度: %d)"+global.GREEN+" [已解决]"+global.RESET, i+1, global.Problems[i].Url, global.Problems[i].Difficulty))
			SendMessage(conn, ais.CommonAI("你是一个温柔的猫娘，性格可爱，语气温柔体贴", global.DARKWHITE+fmt.Sprintf("%d. %s (难度: %d)"+global.GREEN+" [已解决]"+global.RESET, i+1, global.Problems[i].Url, global.Problems[i].Difficulty)))
		}
	}
	SendMessage(conn, fmt.Sprintf("游戏时间: (%d/%d)", (time.Now().Unix()-global.BeginTime)/60, global.GameTime/60))
}

func ListScoresCommand(conn net.Conn) {
	SendMessage(conn, "\n[贡献分列表]:")
	for conn1, score := range global.Score {
		SendMessage(conn, fmt.Sprintf("%s : %d", global.UserNameDisplay[conn1], score))
	}
}
