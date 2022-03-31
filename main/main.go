package main

import (
	"GPA_calculator/ALGO"
	excel "GPA_calculator/excel-process"
	"fmt"
)

func main() {
	//读取excel中的数据，第一列为成绩，第二列为该成绩对应的学分
	data, err := excel.Read("data/Transcript.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	//计算gpa
	GPA, err := ALGO.ToGpa(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(GPA)
}
