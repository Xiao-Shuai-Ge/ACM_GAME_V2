package util

import "ACM_GAME_V2/log/xlog"

func HandleError(err error, text string) {
	xlog.Error(text+"%s", err)
}
