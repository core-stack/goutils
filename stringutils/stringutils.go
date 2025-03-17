package stringutils

import (
	"encoding/json"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func Stringify(v any, indent bool) string {
	if v == nil {
		return ""
	}

	if indent {
		stringValue, _ := json.MarshalIndent(v, "", "  ")
		return string(stringValue)
	}

	stringValue, _ := json.Marshal(v)
	return string(stringValue)
}
func ParseString[T any](s string, target *T) error {
	err := json.Unmarshal([]byte(s), &target)
	return err
}
