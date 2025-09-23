package handler

import (
	"ACM_GAME_V2/global"
	"fmt"
)

func HandleStart() {
	var input string
	// 等待用户输入任意键
	fmt.Println("\n服务端输入任意键开始游戏")
	fmt.Scanln(&input)
	// 模式选择
	fmt.Println("\n模式选择：")
	fmt.Println("1. 简单模式 (6题 60分钟)")
	fmt.Println("2. 简单团战模式 (10题 100分钟)")
	for {
		fmt.Scan(&input)
		switch input {
		case "0":
			global.GameMode = 0
			global.GameTime = 10 * 60
		case "1":
			global.GameMode = 1
			global.GameTime = 60 * 60
		case "2":
			global.GameMode = 2
			global.GameTime = 100 * 60
		default:
			fmt.Println("输入错误")
			continue
		}
		break
	}

}
