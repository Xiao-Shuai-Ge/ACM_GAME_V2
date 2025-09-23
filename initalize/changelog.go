package initalize

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/log/xlog"
	"bufio"
	"os"
	"strings"
)

func InitReadChangelog() {
	fileName := "CHANGELOG.md"
	f, err := os.Open(fileName)
	if err != nil {
		xlog.Error("打开更新日志文件失败: %s", fileName)
		return
	}
	defer f.Close()

	var currentBlock strings.Builder
	var isReading bool
	var cnt int

	global.ChangeLog += "--------------------------------------\n"
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// 开始新块
		if strings.HasPrefix(line, "#") && !isReading {
			if cnt >= 3 {
				break
			}
			if currentBlock.Len() > 0 {
				global.ChangeLog += currentBlock.String()
				currentBlock.Reset()
			}
			isReading = true
		}
		// 结束当前块
		if strings.TrimSpace(line) == "---" && isReading {
			//currentBlock.WriteString(line + "\n")
			global.ChangeLog += currentBlock.String()
			global.ChangeLog += "--------------------------------------\n"
			currentBlock.Reset()
			isReading = false
			cnt++
			continue
		}
		// 累积块内容
		if isReading {
			currentBlock.WriteString(line + "\n")
		}
	}

	// 如果当前块未结束且已有结果少于3，则添加最后一个块
	if isReading && len(global.ChangeLog) < 3 && currentBlock.Len() > 0 {
		global.ChangeLog += currentBlock.String()
	}

	if err != nil {
		xlog.Error("读取更新日志失败: %v", err)
		return
	}

	xlog.Info("解析更新日志成功")
}
