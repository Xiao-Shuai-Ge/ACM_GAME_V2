package handler

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/log/xlog"
	"net"
)

func SendBroadcast(text string) {
	msg := global.Message{
		Text:  text,
		Conn:  nil,
		IsAll: true,
	}

	global.MsgChan <- msg
	xlog.Debug("成功发送一条群发消息: %s", text)
}

func SendMessage(conn net.Conn, text string) {
	msg := global.Message{
		Text:  text,
		Conn:  conn,
		IsAll: false,
	}

	global.MsgChan <- msg
	xlog.Debug("成功对发送一条个人消息: %s", text)
}
