package initalize

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/log/xlog"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	var err error
	global.DB, err = gorm.Open(mysql.Open(global.DSN), &gorm.Config{})
	if err != nil {
		xlog.Error("数据库加载失败")
		return
	}
	xlog.Info("数据库加载成功")
}
