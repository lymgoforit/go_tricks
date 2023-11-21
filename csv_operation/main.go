package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func CreateCsv() {
	//创建文件
	f, err := os.Create("test.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	//创建一个新的写入文件流
	w := csv.NewWriter(f)
	//这个二维数组，外层几个元素就代表几行，内层几个元素就代表几列
	data := [][]string{
		{"序号", "姓名", "年龄"}, //三列
		{"1", "刘备", "23"},
		{"2", "张飞", "23"},
		{"3", "关羽", "23"},
		{"4", "赵云", "23"},
		{"5", "黄忠", "23"},
		{"6", "马超", "23"},
	} //6行
	//写入数据
	w.WriteAll(data)
	w.Flush()
}

func ReadCsv(filePath string) {
	byteContent, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	csvR := csv.NewReader(strings.NewReader(string(byteContent)))
	records, err := csvR.ReadAll()

	if err != nil {
		panic(err)
	}

	fmt.Println(len(records[0][0])) // "\xEF\xBB\xBF" 3字节  + 序号 6字节 = 9字节

	// 针对大文件，一行一行的读取文件
	for i, record := range records {
		if i == 0 {
			continue // 跳过表头
		}
		fmt.Println(record)
	}

}

func main() {
	CreateCsv()
	ReadCsv("test.csv")
}
