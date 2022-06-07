package main

import (
	"bufio"
	data "convert/v1/data"
	"encoding/csv"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/ini.v1"
)

var er string

func main() {
	files, colName := setting()

	for _, file := range files {
		startTime := time.Now()
		t := strings.Split(file.Name(), ".csv")

		if len(t) != 1 {
			fmt.Println(file.Name(), " 파일 필터링을 시작합니다.")
			//Convert
			colName = convertFunc(file.Name(), colName)
			//Time Check
			elapsedTime := time.Since(startTime)
			fmt.Printf("소요 시간: %s \r\n \r\n", elapsedTime)
		}
	}
	fmt.Println("Complete to Filter. Please press Enter")
	fmt.Scanln(&er)
	panic("")
}

func readDirFiles() []fs.FileInfo {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, _ := ioutil.ReadDir(currentDir)
	return files
}

func setting() ([]fs.FileInfo, string) {
	cfg, err := ini.Load("setting.ini")
	if err != nil {
		fmt.Println("Error No setting.ini")
		fmt.Scan(&er)
		panic(err)
	}

	files := readDirFiles()
	col := cfg.Section("SelectCol").Key("Name").String()
	fmt.Printf("Setting.ini Info \r\n Colname :%s \r\n", col)

	return files, col
}

func convertFunc(fileName string, colName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	if colName == "Index" {
		fmt.Println("Index는 필터링할 수 없습니다.")
		fmt.Scan(&er)
		panic("")
	}
	csvReader := csv.NewReader(bufio.NewReader(file))
	rows, _ := csvReader.ReadAll()
	colIdx, colName := data.FindConvertColumn(rows, colName)

	data.FilterBy(rows, colIdx, fileName)

	return colName
}


