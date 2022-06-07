package data

import "strconv"

func ResetIndex(rawData [][]string) [][]string {
	for i := 1; i < len(rawData); i++ {
		rawData[i][0] = strconv.Itoa(i - 1)
	}
	return rawData
}
