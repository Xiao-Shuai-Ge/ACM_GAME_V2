package router

import (
	"ACM_GAME_V2/global"
	"fmt"
	"sync"
)

var (
	mutex = &sync.Mutex{}
)

// HandleMessages 处理消息
func HandleMessages() {
	for {
		message := <-global.MsgChan
		mutex.Lock()
		if message.IsAll {
			// 群发消息
			for conn := range global.Clients {
				_, err := fmt.Fprintf(conn, message.Text+"\n")
				if err != nil {
					fmt.Println("群发信息失败:", err)
				}
			}
		} else {
			// 单发消息
			_, err := fmt.Fprintf(message.Conn, message.Text+"\n")
			if err != nil {
				fmt.Println("发送信息失败:", err)
			}
		}
		mutex.Unlock()
	}
}
