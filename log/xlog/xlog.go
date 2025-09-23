package xlog

import (
	"ACM_GAME_V2/global"
	"fmt"
	"os"
	"time"
)

const (
	DEBUG_LEVEL = 1
	INFO_LEVEL  = 2
	WARN_LEVEL  = 3
	ERROR_LEVEL = 4
)

type xLogger struct {
	level int // 等级
}

var (
	logger xLogger
	ch     chan string
	file   *os.File
)

func InitXLogger(level int) {
	// 初始化日志等级
	logger.level = level
	// 初始化日志保存通道
	ch = make(chan string, 100)
	// 打开日志文件
	dir := "log/data"
	err := os.MkdirAll(dir, 0755) // 如果文件夹不存在，就创建
	if err != nil {
		Error("create directory %s error: %v", dir, err)
		return
	}
	// 日志文件名格式：年-月-日_时-分-秒.log
	fileName := fmt.Sprintf("%s/%s.log", dir, time.Now().Format("2006-01-02_15-04-05"))
	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Error("open file %s error: %v", fileName, err)
		return
	}
	// 启动日志保存协程
	go SaveLog()
}

func SaveLog() {
	defer file.Close()
	defer fmt.Println("log save goroutine exit")
	for {
		msg := <-ch
		_, err := file.WriteString(msg)
		if err != nil {
			Error("write file error: %v", err)
			return
		}
	}
}

func Debug(format string, args ...interface{}) {
	if logger.level <= DEBUG_LEVEL {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+global.CYAN+" [DEBUG]"+global.RESET+" "+format+"\n", args...)
		ch <- fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")+" [DEBUG] "+format+"\n", args...)
	}
}

func Info(format string, args ...interface{}) {
	if logger.level <= INFO_LEVEL {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+global.GREEN+" [INFO]"+global.RESET+" "+format+"\n", args...)
		ch <- fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")+" [INFO] "+format+"\n", args...)
	}
}

func Warn(format string, args ...interface{}) {
	if logger.level <= WARN_LEVEL {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+global.YELLOW+" [WARN]"+global.RESET+" "+format+"\n", args...)
		ch <- fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")+" [WARN] "+format+"\n", args...)
	}
}

func Error(format string, args ...interface{}) {
	if logger.level <= ERROR_LEVEL {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+global.RED+" [ERROR]"+global.RESET+" "+format+"\n", args...)
		ch <- fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")+" [ERROR] "+format+"\n", args...)
	}
}
