package text

import (
	"fmt"
	"strings"
)

func SimpleNormalize(text string) string {
	return strings.Join(strings.Fields(text), " ")
}

func CheckResult(output []string) bool {
	if len(output) == 0 {
		fmt.Println("По вашему запросу ничего не было найдено(")
		return false
	}
	return true
}
