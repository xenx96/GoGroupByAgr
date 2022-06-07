package data

import (
	"fmt"
	"strconv"
	"strings"
)

func FilterBy(data [][]string, colNum int, fileName string){
	var uniqueGroup []string

	for i, row := range data {
		if i == 0 {
			continue
		}
		uniqueGroup = append(uniqueGroup, row[colNum])
	}

	uniqueGroup = unique(uniqueGroup)
	i := 0
	for _, group := range uniqueGroup {
		result := filterArray(data, group, colNum)
		
		fileSplit := strings.Split(fileName,".csv")
		saveName := fileSplit[0]+"_"+group+".csv"
		WriteCSV(result, saveName)
		
		fmt.Println(saveName+"으로 필터링되었습니다.")
		
		i++
	}
}

func filterArray(data [][]string, group string, colNum int) [][]string {
	var result [][]string
	result = append(result, data[0])
	gr, _ := strconv.Atoi(group)
	for _, row := range data {
		i, err := strconv.Atoi(row[colNum])
		if err != nil {
			continue
		}

		if i == gr {
			result = append(result, row)
		}
	}

	return result
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
