package handler

import (
	"ACM_GAME_V2/common/ais"
	"ACM_GAME_V2/global"
	"ACM_GAME_V2/internal/repo"
	"ACM_GAME_V2/log/xlog"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

// Codeforces API 响应结构体
type CodeforcesResponse struct {
	Status  string      `json:"status"`
	Comment string      `json:"comment"`
	Result  interface{} `json:"result"`
}

// 提交记录结构体
type CodeforcesSubmission struct {
	ID                  int     `json:"id"`
	ContestID           int     `json:"contestId"`
	Problem             Problem `json:"problem"`
	Author              Author  `json:"author"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict             string  `json:"verdict"`
	Testset             string  `json:"testset"`
	PassedTestCount     int     `json:"passedTestCount"`
	TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
	CreationTimeSeconds int64   `json:"creationTimeSeconds"`
}

type Problem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Points    float64  `json:"points"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}

type Author struct {
	ContestID        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	Ghost            bool     `json:"ghost"`
	StartTimeSeconds int64    `json:"startTimeSeconds"`
}

type Member struct {
	Handle string `json:"handle"`
}

// 获取用户提交记录的API函数
func GetUserSubmissionsFromAPI(username string) ([]CodeforcesSubmission, error) {
	// 添加随机延迟避免频率限制
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+500))

	// 构建API请求URL
	url := fmt.Sprintf("https://codeforces.com/api/user.status?handle=%s&from=1&count=100", username)

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("User-Agent", "ACM_GAME_V2/1.0")
	req.Header.Set("Accept", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API请求失败，状态码: %d", resp.StatusCode)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析JSON响应
	var apiResponse CodeforcesResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	// 检查API响应状态
	if apiResponse.Status != "OK" {
		return nil, fmt.Errorf("API返回错误: %s", apiResponse.Comment)
	}

	// 将result转换为提交记录数组
	resultBytes, err := json.Marshal(apiResponse.Result)
	if err != nil {
		return nil, fmt.Errorf("序列化result失败: %v", err)
	}

	var submissions []CodeforcesSubmission
	if err := json.Unmarshal(resultBytes, &submissions); err != nil {
		return nil, fmt.Errorf("解析提交记录失败: %v", err)
	}

	return submissions, nil
}

func Check() {
	SendBroadcast(global.SEND_XITONG + " 进行一次提交检查...")
	xlog.Info("进行一次检查")
	global.CheckTime = time.Now().Unix()
	GetSubmission()
	for i := 0; i < len(global.Problems); i++ {
		if global.ProblemPenalty[i] != -1 {
			continue
		}
		AC := false
		WA := 0
		var AC_User net.Conn
		for conn := range global.Clients {
			ac, wa := CheckProblem(conn, global.Problems[i].Url, i+1)
			if ac && !AC {
				AC = true
				AC_User = conn
			}
			WA += wa
		}
		if AC {
			global.ProblemPenalty[i] = int(time.Now().Unix()-global.BeginTime)/60 + WA*10
			SendBroadcast(fmt.Sprintf(global.SEND_XITONG+" %s "+global.GREEN+"AC了第 %d 题！(罚时: %d )"+global.RESET, global.UserNameDisplay[AC_User], i+1, global.ProblemPenalty[i]))
			SendBroadcast(ais.CommonAI("你是一个温柔的猫娘，语气可爱，声音甜美", fmt.Sprintf(global.SEND_XITONG+" %s "+global.GREEN+"AC了第 %d 题！(罚时: %d )"+global.RESET, global.UserNameDisplay[AC_User], i+1, global.ProblemPenalty[i])))
			SendMessage(AC_User, fmt.Sprintf(global.YELLOW+" + %d 贡献分 (参加游戏得分)"+global.RESET, global.Problems[i].Difficulty/10))
			global.Score[AC_User] += global.Problems[i].Difficulty / 10
			repo.MarkUsed(global.Problems[i].Url)
		}
	}
}

func CheckProblem(conn net.Conn, Url string, id int) (ac bool, wa int) {
	for _, submission := range global.UserSubmission[conn] {
		if submission.Url == Url {
			if submission.Status == "Accepted" {
				ac = true
			} else if submission.Status == "Compilation error" {
			} else if len(submission.Status) >= 4 && (submission.Status[:4] == "Runn" || submission.Status[:4] == "In q") {
			} else {
				wa += 1
				if !global.WA_ID[submission.SubmissionID] {
					global.WA_ID[submission.SubmissionID] = true
					SendBroadcast(fmt.Sprintf(global.SEND_XITONG+" %s "+global.RED+"WA了第 %d 题！"+global.RESET, global.UserName[conn], id))
				}
			}
		}
	}
	return ac, wa
}

func GetSubmission() {
	for conn := range global.Clients {
		username := global.UserName[conn]
		xlog.Info("正在获取用户%s的提交记录", username)

		// 使用Codeforces API获取提交记录
		submissions, err := GetUserSubmissionsFromAPI(username)
		if err != nil {
			xlog.Error("获取用户%s的提交记录失败: %v", username, err)
			continue
		}

		xlog.Info("获取用户%s的提交记录成功，共%d条记录", username, len(submissions))

		// 转换API数据格式为内部格式
		global.UserSubmission[conn] = make([]global.Submission, 0)

		for _, submission := range submissions {
			// 构建题目URL
			problemUrl := fmt.Sprintf("https://codeforces.com/problemset/problem/%d/%s",
				submission.Problem.ContestID, submission.Problem.Index)

			// 转换提交状态
			status := convertVerdictToStatus(submission.Verdict)

			// 添加到用户提交记录
			global.UserSubmission[conn] = append(global.UserSubmission[conn], global.Submission{
				SubmissionID: strconv.Itoa(submission.ID),
				Url:          problemUrl,
				Status:       status,
			})
		}
	}
}

// 转换Codeforces verdict到内部状态格式
func convertVerdictToStatus(verdict string) string {
	switch verdict {
	case "OK":
		return "Accepted"
	case "WRONG_ANSWER":
		return "Wrong answer"
	case "TIME_LIMIT_EXCEEDED":
		return "Time limit exceeded"
	case "MEMORY_LIMIT_EXCEEDED":
		return "Memory limit exceeded"
	case "RUNTIME_ERROR":
		return "Runtime error"
	case "COMPILATION_ERROR":
		return "Compilation error"
	case "PRESENTATION_ERROR":
		return "Presentation error"
	case "IDLENESS_LIMIT_EXCEEDED":
		return "Idleness limit exceeded"
	case "SECURITY_VIOLATED":
		return "Security violated"
	case "CRASHED":
		return "Crashed"
	case "INPUT_PREPARATION_CRASHED":
		return "Input preparation crashed"
	case "CHALLENGED":
		return "Challenged"
	case "SKIPPED":
		return "Skipped"
	case "TESTING":
		return "Running"
	case "REJECTED":
		return "Rejected"
	default:
		return verdict
	}
}
