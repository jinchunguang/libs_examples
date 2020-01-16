package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"time"
)

func main() {

	t := time.Now()

	sheetName := "Sheet1"

	f := excelize.NewFile()
	// 创建一个新的sheet
	index := f.NewSheet(sheetName)

	// 设置数据
	for i := 1; i <= 30000; i++ {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(i), "PHP."+strconv.Itoa(i))
		f.SetCellValue(sheetName, "B"+strconv.Itoa(i), "JAVA."+strconv.Itoa(i))
		f.SetCellValue(sheetName, "C"+strconv.Itoa(i), "C++."+strconv.Itoa(i))
		f.SetCellValue(sheetName, "D"+strconv.Itoa(i), "PYTHON."+strconv.Itoa(i))
	}

	// 保存文件到指定目录
	f.SetActiveSheet(index)
	err := f.SaveAs("./test.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	elapsed := time.Since(t)
	fmt.Println("Excel写入:", 30000, " 耗时:", elapsed)
}
