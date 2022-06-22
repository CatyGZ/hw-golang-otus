package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type FrequencyAnalytic struct {
	Word      string
	Frequency int
}

const cntResultValue = 10

var regexExt = regexp.MustCompile(`[\!\,\.\?\:\;]`)

func Top10(sourceStr string) []string {
	if sourceStr == "" {
		return nil
	}

	resultArray := make([]string, 0)
	frequencyArray := make([]FrequencyAnalytic, 0)

	sourceArray := strings.Fields(sourceStr)
	for _, v := range sourceArray {
		word := regexExt.ReplaceAllString(strings.ToLower(v), "")
		if word == "-" {
			continue
		}

		// проверка - встречалось ли слово ранее
		ind, ok := contains(frequencyArray, word)
		if !ok {
			frequencyArray = append(frequencyArray, FrequencyAnalytic{Word: word, Frequency: 1})
		} else {
			frequencyArray[ind].Frequency++
		}
	}

	// сортировка
	sort.Slice(frequencyArray, func(i, j int) bool {
		if frequencyArray[i].Frequency == frequencyArray[j].Frequency {
			return frequencyArray[i].Word < frequencyArray[j].Word
		}
		return frequencyArray[i].Frequency > frequencyArray[j].Frequency
	})

	// формирование результирующего массива
	for i, v := range frequencyArray {
		resultArray = append(resultArray, v.Word)
		if i >= cntResultValue-1 {
			break
		}
	}

	return resultArray
}

// contains указывает, содержится ли строка в массиве строк.
func contains(array []FrequencyAnalytic, searchStr string) (int, bool) {
	for i, v := range array {
		if strings.EqualFold(v.Word, searchStr) {
			return i, true
		}
	}
	return 0, false
}
