package data

import (
	"math"
	"strconv"
)

func noValue() float64 {
	return -99999999999
}

// colNum == timeColIdx
func GroupBy(data [][]string, colNum int, stats int64) [][]string {
	var result [][]string
	var uniqueGroup []string

	for i, row := range data {
		if i == 0 {
			result = append(result, row)
			continue
		}
		uniqueGroup = append(uniqueGroup, row[colNum])
	}

	uniqueGroup = unique(uniqueGroup)
	i := 0
	for _, group := range uniqueGroup {
		result = append(result, checkGroup(data, group, colNum, stats))
		i++
	}

	return result
}

func checkGroup(data [][]string, group string, colNum int, stats int64) []string {
	var result [][]string
	ary := make([]string, len(data[0]))

	gr, _ := strconv.Atoi(group)
	for _, row := range data {
		i, err := strconv.Atoi(row[colNum])
		if err != nil {
			continue
		}

		if i == gr {
			result = append(result, row)
		} else if i < gr {
			continue
		} else if i > gr {
			break
		}
	}

	ary[0] = group
	col := data[0]
	if stats > 3 {
		return findWhere(result, col, ary, 0, stats)
	}
	_, ary = aggregation(result, col, ary, 0, stats)
	return ary
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

func aggregation(groupData [][]string, col []string, result []string, w int, stats int64) (int, []string) {
	var ary []string
	min := noValue()
	max := noValue()
	val := ""

	if w == (len(col)) {
		return w, result
	}

	for _, data := range groupData {
		value, err := strconv.ParseFloat(data[w], 64)
		if err != nil {
			continue
		}

		if min == noValue() || max == noValue() {
			min = value
			max = value
			ary = append(ary, data[w])
			continue
		}

		min = math.Min(min, value)
		max = math.Max(max, value)
		ary = append(ary, data[w])
	}

	switch stats {
	case 1:
		val = strconv.FormatFloat(min, 'f', -1, 32)
	case 2:
		val = strconv.FormatFloat(max, 'f', -1, 32)
	case 3:
		val = avgFunc(ary)
	}
	if min == noValue() || max == noValue() {
		val = ""
	}

	result[w] = val
	return aggregation(groupData, col, result, w+1, stats)
}

func avgFunc(ary []string) string {
	sum := 0.0
	n := len(ary)
	if n == 1 {
		return ary[0]
	}

	for i := 0; i < n; i++ {
		f, _ := strconv.ParseFloat(ary[i], 64)
		sum += f
	}

	return strconv.FormatFloat(sum/(float64(n)), 'f', -1, 32)
}

func findWhere(groupData [][]string, col []string, result []string, w int, stats int64)[]string{
	where := 0
	if w == (len(col)) {
		return result
	}
	if len(groupData)== 1{
		where = 0
	}else if stats==6 {
		where = len(groupData)-1
	}else if stats ==5{
		where = len(groupData)/2
	}

	data := groupData[where]
	result[w] = data[w]
	return findWhere(groupData, col, result, w+1, stats)
}