package util

import "ACM_GAME_V2/global"

func ToUserNameDisplay(score int, username string) string {
	if score <= 200 {
		return global.DARKWHITE + username + global.RESET
	} else if score <= 500 {
		return global.GREEN + username + global.RESET
	} else if score <= 800 {
		return global.CYAN + username + global.RESET
	} else if score <= 1200 {
		return global.BLUE + username + global.RESET
	} else if score <= 1600 {
		return global.MAGENTA + username + global.RESET
	} else if score <= 2000 {
		return global.YELLOW + username + global.RESET
	} else {
		return global.RED + username + global.RESET
	}
}
