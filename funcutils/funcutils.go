package funcutils

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFunctionName(i any) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	lastIndex := strings.LastIndex(fullName, ".")
	if lastIndex == -1 {
		return fullName
	}
	return fullName[lastIndex+1:]
}
