package excel_process

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

//从excel中读取自己的学分和成绩，用的是开源的库excelize，返回二维数组
//这里不好的是，我没有返回一个切块，导致必须在excel中查看有多少们课，这里是63
//后续可以改进

func Read(path string) ([63][2]float64, error) {
	var numbers [63][2]float64
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return numbers, nil
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//get all cells from excel
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return numbers, err
	}
	numbers = convert(rows)
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}
	return numbers, nil
}

//转换将读取到的数据由string转换为float64类型的二维数组数据
func convert(rows [][]string) [63][2]float64 {
	var numbers [63][2]float64
	for i, row := range rows {
		numbers[i][0], _ = strconv.ParseFloat(row[0], 64)
		numbers[i][1], _ = strconv.ParseFloat(row[1], 64)
	}
	return numbers
}
