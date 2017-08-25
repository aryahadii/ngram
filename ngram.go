package ngram

import (
	"errors"
)

// Get creates a string array contains ngram
func Get(text string, n int) ([]string, error) {
	runeArr := []rune(text)
	result := []string{}

	if n < 1 || n > len(runeArr) {
		return result, errors.New(`"n" is more than "text" length`)
	}

	for index := 0; index < len(runeArr)-n+1; index++ {
		result = append(result, string(runeArr[index:index+n]))
	}

	return result, nil
}

// GetFrequency create a map which contains frequency of each rune groups
func GetFrequency(text string, n int) (map[string]int, error) {
	runeArr := []rune(text)
	result := make(map[string]int)

	if n < 1 || n > len(runeArr) {
		return result, errors.New(`"n" is more than "text" length`)
	}

	for index := 0; index < len(runeArr)-n+1; index++ {
		key := string(runeArr[index : index+n])
		if _, ok := result[key]; ok {
			result[key]++
		} else {
			result[key] = 1
		}
	}

	return result, nil

}
