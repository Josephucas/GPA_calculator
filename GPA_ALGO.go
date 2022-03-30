package GPA_ALGO

import (
	"errors"
)

func ToGp(score int) (float64, error) {
	var Gp float64
	switch {
	case 100 < score:
		return 4, errors.New("输入的值大于100，出现错误！！")
	case 90 <= score:
		return 4, nil
	case score < 60:
		return 0, nil
	default:
		Gp = float64(score-50) / 10
		return Gp, nil
	}
}
