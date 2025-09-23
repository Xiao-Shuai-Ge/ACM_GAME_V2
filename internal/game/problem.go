package game

import (
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/internal/handler"
	"ACM_GAME_V2/internal/repo"
	"ACM_GAME_V2/log/xlog"
	"errors"
	"math/rand"
)

func GetProblem() {
	// 先获取题目
	var err error
	if global.GameMode == 0 {
		err = GetProblemTest()
	} else if global.GameMode == 1 {
		err = GetProblemEasy()
	} else if global.GameMode == 2 {
		err = GetProblemEasyTeam()
	}
	if err != nil {
		xlog.Error("生成题目失败: %v", err)
		return
	}
	// 将罚时设置为-1，代表没有解决此题
	for i := 0; i < len(global.Problems); i++ {
		global.ProblemPenalty[i] = -1
	}
	// 随机排序
	RandomSort()
}

func RandomSort() {
	for i := 0; i < 100; i++ {
		l := rand.Intn(len(global.Problems))
		r := rand.Intn(len(global.Problems))
		global.Problems[l], global.Problems[r] = global.Problems[r], global.Problems[l]
	}
}

func GetProblemTest() error {
	ProblemNum := 1
	global.Problems = make([]global.Problem, ProblemNum)
	global.ProblemPenalty = make([]int, ProblemNum)

	global.Problems[0].Url = "https://codeforces.com/problemset/problem/1981/B"
	global.Problems[0].Difficulty = 1300

	handler.SendBroadcast(`
    _    ___  __  __         _____  ___  ___  _____   __  __   ___   ___   ___ 
   /_\  / __||  \/  |  ___  |_   _|| __|/ __||_   _| |  \/  | / _ \ |   \ | __|
  / _ \| (__ | |\/| | |___|   | |  | _| \__ \  | |   | |\/| || (_) || |) || _| 
 /_/ \_\\___||_|  |_|         |_|  |___||___/  |_|   |_|  |_| \___/ |___/ |___|
`)
	return nil
}

func GetProblemEasy() error {
	ProblemNum := 6
	global.Problems = make([]global.Problem, ProblemNum)
	global.ProblemPenalty = make([]int, ProblemNum)
	var err1, err2, err3, err4, err5, err6 error
	global.Problems[0], err1 = repo.GetProblem(800, 800)
	global.Problems[1], err2 = repo.GetProblem(800, 900)
	global.Problems[2], err3 = repo.GetProblem(900, 1000)
	global.Problems[3], err4 = repo.GetProblem(1000, 1200)
	global.Problems[4], err5 = repo.GetProblem(1200, 1500)
	global.Problems[5], err6 = repo.GetProblem(1500, 1800)
	combinedErr := errors.Join(err1, err2, err3, err4, err5, err6)

	handler.SendBroadcast(`
    _    ___  __  __         ___    _    ___ __   __  __  __   ___   ___   ___ 
   /_\  / __||  \/  |  ___  | __|  /_\  / __|\ \ / / |  \/  | / _ \ |   \ | __|
  / _ \| (__ | |\/| | |___| | _|  / _ \ \__ \ \ V /  | |\/| || (_) || |) || _| 
 /_/ \_\\___||_|  |_|       |___|/_/ \_\|___/  |_|   |_|  |_| \___/ |___/ |___|
`)

	return combinedErr
}

func GetProblemEasyTeam() error {
	ProblemNum := 10
	global.Problems = make([]global.Problem, ProblemNum)
	global.ProblemPenalty = make([]int, ProblemNum)
	var err1, err2, err3, err4, err5, err6, err7, err8, err9, err10 error
	global.Problems[0], err1 = repo.GetProblem(800, 800)
	global.Problems[1], err2 = repo.GetProblem(800, 800)
	global.Problems[2], err3 = repo.GetProblem(800, 900)
	global.Problems[3], err4 = repo.GetProblem(800, 900)
	global.Problems[4], err5 = repo.GetProblem(900, 1000)
	global.Problems[5], err6 = repo.GetProblem(1000, 1200)
	global.Problems[6], err7 = repo.GetProblem(1100, 1400)
	global.Problems[7], err8 = repo.GetProblem(1200, 1500)
	global.Problems[8], err9 = repo.GetProblem(1500, 1800)
	global.Problems[9], err10 = repo.GetProblem(1700, 2000)
	combinedErr := errors.Join(err1, err2, err3, err4, err5, err6, err7, err8, err9, err10)

	handler.SendBroadcast(`
    _    ___  __  __         ___    _    ___ __   __  _____  ___    _    __  __   __  __   ___   ___   ___ 
   /_\  / __||  \/  |  ___  | __|  /_\  / __|\ \ / / |_   _|| __|  /_\  |  \/  | |  \/  | / _ \ |   \ | __|
  / _ \| (__ | |\/| | |___| | _|  / _ \ \__ \ \ V /    | |  | _|  / _ \ | |\/| | | |\/| || (_) || |) || _| 
 /_/ \_\\___||_|  |_|       |___|/_/ \_\|___/  |_|     |_|  |___|/_/ \_\|_|  |_| |_|  |_| \___/ |___/ |___|
                                                                                                           
`)

	return combinedErr
}

func ListProblemsAll() {
	for conn := range global.Clients {
		handler.ListProblemsCommand(conn)
	}
}
