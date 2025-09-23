package repo

import (
	"ACM_GAME_V2/global"
)

type Problem struct {
	Url        string `gorm:"type:varchar(100);null;"`
	Difficulty int    `gorm:"type:int;null;"`
	IsUsed     bool   `gorm:"type:tinyint(1);null;"`
}

func GetProblem(difficulty_l int, difficulty_r int) (problem global.Problem, err error) {
	//fmt.Println("difficulty_l:", difficulty_l, "difficulty_r:", difficulty_r)
	err = global.DB.Where("difficulty >= ? AND difficulty <= ? AND is_used = 0", difficulty_l, difficulty_r).Order("RAND()").First(&problem).Error
	if err != nil {
		return
	}
	for i := 0; i < len(global.Problems); i++ {
		if global.Problems[i].Url == problem.Url {
			problem, err = GetProblem(difficulty_l, difficulty_r)
		}
	}
	//err = global.DB.Model(&Problem{}).Where("url = ?", problem.Url).Update("is_used", 1).Error
	return
}

func MarkUsed(url string) (err error) {
	err = global.DB.Model(&Problem{}).Where("url = ?", url).Update("is_used", 1).Error
	return
}
