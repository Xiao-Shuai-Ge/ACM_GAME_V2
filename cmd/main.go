package main

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/initalize"
	"ACM_GAME_V2/internal/game"
	"ACM_GAME_V2/internal/handler"
	"ACM_GAME_V2/internal/repo"
	"ACM_GAME_V2/internal/router"
	"ACM_GAME_V2/log/xlog"
	"fmt"
	"net"
	"time"
)

func main() {
	// 初始化
	initalize.Init()
	xlog.Info("ACM GAME 启动成功!")
	xlog.Info("目前版本: %s", global.Version)

	// 开启服务端
	var err error
	global.Listener, err = net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("开启服务端失败:", err)
		return
	}
	defer global.Listener.Close()

	//开启公开聊天服务
	go router.HandleMessages()
	go router.Join()

	// 更新分数
	repo.UpdateScore()

	for {
		handler.HandleStart()
		game.Game()
	}

	time.Sleep(100 * time.Second)
}
