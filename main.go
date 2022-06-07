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
	"strconv"
	"strings"
	"time"

	"gopkg.in/ini.v1"
)

var er string

func main() {
	files, time_col, stats, convertValue := setting()

	for _, file := range files {
		startTime := time.Now()
		t := strings.Split(file.Name(), ".csv")

		if len(t) != 1 {
			fmt.Println(file.Name(), " 파일 변경을 시작합니다.")
			//Convert
			time_col = convertFunc(file.Name(), time_col, stats, convertValue)
			//Time Check
			elapsedTime := time.Since(startTime)
			fmt.Printf("소요 시간: %s\n", elapsedTime)
		}
	}
	fmt.Println("Complete to Convert. Please press Enter")
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

func setting() ([]fs.FileInfo, string, int64, int64) {
	cfg, err := ini.Load("setting.ini")
	var s int64
	if err != nil {
		fmt.Println("Error No setting.ini")
		fmt.Scan(&er)
		panic(err)
	}

	files := readDirFiles()
	col := cfg.Section("SelectCol").Key("Name").String()
	nCompression, _ := cfg.Section("Options").Key("N_Compression").Int()
	stat := cfg.Section("Options").Key("Stat").String()
	fmt.Printf("Setting.ini Info \r\n Colname :%s \r\n Stat : %s \r\n 압축비율 : 1/%s \r\n \r\n ", col, stat, strconv.Itoa(nCompression))
	if stat == "min" {
		s = 1
	} else if stat == "max" {
		s = 2
	} else if stat == "avg" {
		s = 3
	} else if stat == "first" {
		s = 4
	} else if stat == "center" {
		s = 5
	} else if stat == "last" {
		s = 6
	} else {
		fmt.Println("Error! Stat 정보가 잘못 되었습니다.")
		fmt.Scan(&er)
		panic("")
	}
	return files, col, s, int64(nCompression)
}

func convertFunc(fileName string, timeCol string, stats int64, convertValue int64) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(bufio.NewReader(file))
	rows, _ := csvReader.ReadAll()
	timeColIdx, timeCol := data.FindConvertColumn(rows, timeCol)

	i := -1
	for _, row := range rows {
		i++
		var index []string
		if i == 0 {
			index = append(index, "Index")
			rows[i] = append(index, row...)
			continue
		}
		index = append(index, strconv.Itoa(i))
		row = append(index, row...)
		row[timeColIdx] = secondConvert(row[timeColIdx], convertValue)
		rows[i] = row
	}

	result := data.GroupBy(rows, timeColIdx, stats)
	if timeCol != "Index" {
		result = data.ResetIndex(result)
	}
	data.WriteCSV(result, fileName)
	return timeCol
}

func secondConvert(row string, convertValue int64) string {
	time, _ := strconv.Atoi(row)
	row = strconv.Itoa(time / int(convertValue))
	return row
}

