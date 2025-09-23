package router

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/internal/handler"
	"ACM_GAME_V2/internal/repo"
	"ACM_GAME_V2/log/xlog"
	"ACM_GAME_V2/util"
	"bufio"
	"fmt"
	"net"
)

func Join() {
	for {
		conn, err := global.Listener.Accept()
		if err != nil {
			fmt.Println("连接错误", err)
			return
		}

		if global.GameState == 1 {
			_, err := fmt.Fprintf(conn, "你中途加入了游戏\n")
			if err != nil {
				fmt.Println("发送信息失败:", err)
			}
		}

		// 标记已连接
		mutex.Lock()
		global.Clients[conn] = true
		mutex.Unlock()

		fmt.Fprintf(conn, "欢迎来到ACM GAME,可以输入/help查看可用指令\n")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		xlog.Info("玩家 %s 离开聊天室", global.UserName[conn])
		handler.SendBroadcast(fmt.Sprintf("["+global.YELLOW+"系统"+global.RESET+"] %s "+global.YELLOW+"离开聊天室"+global.RESET, global.UserName[conn]))
		//关闭连接
		mutex.Lock()
		delete(global.Clients, conn)
		// delete(global.UserName, conn)
		mutex.Unlock()
		err := conn.Close()
		if err != nil {
			return
		}
	}()

	GetInput(conn)
}

func GetInput(conn net.Conn) {
	var getName bool
	//等待获取输入
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		if len(message) == 0 {
			continue
		}
		if getName == false {
			// 处理玩家用户名
			getName = true
			mutex.Lock()
			global.UserName[conn] = message
			global.ConnList[global.UserName[conn]] = conn

			// 数据库处理分数
			score, _ := repo.GetUser(global.UserName[conn])
			global.UserNameDisplay[conn] = util.ToUserNameDisplay(score, global.UserName[conn])

			xlog.Info("玩家 %s 进入聊天室", global.UserName[conn])
			handler.SendBroadcast(fmt.Sprintf("["+global.YELLOW+"系统"+global.RESET+"] %s(%d) "+global.YELLOW+"进入聊天室"+global.RESET, global.UserNameDisplay[conn], score))
			mutex.Unlock()
		} else if message[0] == '/' {
			handler.HandleCommand(conn, message)
		} else {
			handler.SendBroadcast(fmt.Sprintf("[%s]:%s", global.UserNameDisplay[conn], message))
		}
	}
}
