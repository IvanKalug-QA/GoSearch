package text

import "strings"

func SimpleNormalize(text string) string {
	return strings.Join(strings.Fields(text), " ")
}
