package utils

import (
	"fmt"
)

func Assert(isTrue bool, msg string, args ...any) {
	if isTrue {
		return
	}
	panic(fmt.Sprintf(msg, args...))
}
