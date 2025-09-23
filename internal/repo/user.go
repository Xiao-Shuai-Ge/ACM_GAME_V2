package repo

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/log/xlog"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func GetUser(username string) (score int, point int) {
	var cnt int64 // 出现次数，用来检测用户是否存在
	err := global.DB.Model(&global.User{}).Where("user_name =?", username).Count(&cnt)
	if err != nil {
		xlog.Error("查询用户失败:", err)
	}
	if cnt == 0 {
		// 用户不存在
		xlog.Info("用户不存在，创建新用户:", username)
		global.DB.Create(&global.User{UserName: username, Score: 0, Point: 0})
		score = 0
		point = 0
	} else {
		// 用户存在
		var user global.User
		err := global.DB.Model(&global.User{}).Where("user_name =?", username).First(&user).Error
		if err != nil {
			xlog.Error("查询用户失败:", err)
		}
		score = user.Score
		point = user.Point
	}
	return
}

func AddScore(username string, score int) {
	err := global.DB.Model(&global.User{}).Where("user_name =?", username).Update("score", gorm.Expr("score + ?", score)).Error
	if err != nil {
		xlog.Error("更新用户分数失败:", err)
	}
}

func UpdateScore() {
	// 定时更新用户分数
	var cnt int64 // 出现次数，用来检测上次更新时间是否存在
	global.DB.Model(&global.Text{}).Where("k = ?", "last_update_score_time").Count(&cnt)
	if cnt == 0 {
		// 上次更新时间不存在，创建新记录，并且不需要更新用户分数
		global.DB.Create(&global.Text{K: "last_update_score_time", V: fmt.Sprintf("%v", int64(time.Now().Unix())/(60*60*24))})
	} else {
		// 上次更新时间存在，更新用户分数
		var text global.Text
		err := global.DB.Model(&global.Text{}).Where("k = ?", "last_update_score_time").First(&text).Error
		if err != nil {
			xlog.Error("查询上次更新时间失败:", err)
		}
		last_update_time, err := strconv.ParseInt(text.V, 10, 64)
		if err != nil {
			xlog.Error("解析上次更新时间失败:", err)
		}
		diff_day := (time.Now().Unix() / (60 * 60 * 24)) - last_update_time
		if diff_day >= 1 {
			// 超过1天才更新用户分数
			var users []global.User
			err := global.DB.Model(&global.User{}).Find(&users).Error
			if err != nil {
				xlog.Error("查询用户失败:", err)
			}
			for _, user := range users {
				// 计算用户分数
				score := user.Score
				// 计算用户分数
				score -= int(diff_day) * 10 // 每天减少10分
				if score < 0 {
					score = 0
				} else if score > 2500 {
					score = 2500
				}
				// 更新用户分数
				err := global.DB.Model(&global.User{}).Where("user_name =?", user.UserName).Update("score", score).Error
				if err != nil {
					xlog.Error("更新用户分数失败:", err)
				}
			}
			// 更新上次更新时间
			global.DB.Model(&global.Text{}).Where("k = ?", "last_update_score_time").Update("v", fmt.Sprintf("%v", int64(time.Now().Unix())/(60*60*24)))
		}
	}
}
