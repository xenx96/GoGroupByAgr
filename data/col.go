package data

import (
	"fmt"
	"strconv"
)

func FindConvertColumn(data [][]string, col string) (int, string) {
	headers := data[0]
	if col == "Index" {
		return 0, col
	}
	for i, h := range headers {
		if h == col {
			return i + 1, col
		}
	}

	fmt.Println("Error! 해당 Column은 존재하지 않습니다.")
	timeCol := InputColum()
	return FindConvertColumn(data, timeCol)
}

func InputColum() string {
	var time_col string

	fmt.Printf("1. 변경할 시간 Column을 입력해주세요 - 없을 경우 Enter(Index 기준) \r\n(대소문자 구분) : ")
	fmt.Scanln(&time_col)
	if time_col == "" {
		time_col = "Index"
	}
	fmt.Println(time_col + "을 입력받았습니다.\r\n")
	return time_col
}

func cTime() []int64 {
	var ary []int64
	var ms, s, m, h int64
	ms, s, m, h = 1, 1000, 60000, 3600000

	ary = append(ary, ms)
	ary = append(ary, s)
	ary = append(ary, m)
	ary = append(ary, h)

	return ary
}

func InputConvertValue(colname string) int64 {
	var txt string
	convertTime := cTime()

	if colname == "Index" {
		fmt.Printf("2. 압축 비율을 입력해주세요. (1/N) \r\n (input N) : ")
		fmt.Scanln(&txt)
		zipValue, _ := strconv.ParseInt(txt, 10, 64)
		fmt.Println(txt + "을 입력받았습니다.\r\n")
		return zipValue
	}

	fmt.Printf("2-1. 현재 데이터의 시간 단위를 선택해주세요 - (번호 입력) \r\n 1-(Mili-Second) 2-(Second) 3-(Minute) : ")
	fmt.Scanln(&txt)
	firstValue, _ := strconv.ParseInt(txt, 10, 64)
	firstValue = convertTime[firstValue-1]
	fmt.Println(txt + "을 입력받았습니다.\r\n")

	fmt.Printf("2-2. 변경하실 시간 단위를 선택해주세요 - (번호 입력) \r\n 1-(Second) 2-(Minute) 3-(Hour) : ")
	fmt.Scanln(&txt)
	convertValue, _ := strconv.ParseInt(txt, 10, 64)
	convertValue = convertTime[convertValue]
	fmt.Println(txt + "을 입력받았습니다.\r\n")

	if firstValue >= convertValue {
		fmt.Println("시간 단위 선택이 잘못되었습니다. 다시 입력해주세요. \r\n ")
		return InputConvertValue(colname)
	}

	result := convert(firstValue, convertValue)
	return result

}

func convert(beforeTime int64, convertTime int64) int64 {
	time := convertTime / beforeTime
	return time
}

func InputStats() int64 {
	var stats string

	fmt.Printf("3. 무엇을 선택하시겠습니까? - (번호 선택) \r\n1.Min(최소) 2.Max(최대) 3.Avg(평균) : ")
	fmt.Scanln(&stats)
	fmt.Println(stats + "을 입력받았습니다.\r\n")

	if stats == "1" || stats == "2" || stats == "3" {
		result, _ := strconv.ParseInt(stats, 10, 64)
		return result
	} else {
		fmt.Println("잘못된 선택입니다. 다시 선택해주세요 \r\n ")
		return InputStats()
	}

}
