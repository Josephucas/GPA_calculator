package ALGO

import (
	"errors"
)

//百分之学分绩点计算方法
//将单科成绩和学分转化为单门课程的学分绩点
func toGp(score []float64) (float64, error) {
	var Gp float64
	switch {
	case 100 < score[0]:
		return 4 * score[1], errors.New("输入的值大于100，出现错误！！")
	case 90 <= score[0]:
		return 4 * score[1], nil
	case score[0] < 60:
		return 0, nil
	default:
		Gp = float64(score[0]-50) / 10
		return Gp * score[1], nil
	}
}

//计算全部科目的平均学分绩
func ToGpa(rows [][]float64) (float64, error) {

	var Gps, credits float64

	for _, row := range rows {
		gp, _ := toGp(row)
		Gps += gp
		credits += row[1]
	}

	GPA := Gps / credits
	return GPA, nil
}
