package initalize

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/log/xlog"
)

func InitLog() {
	xlog.InitXLogger(global.LogLevel)
}
