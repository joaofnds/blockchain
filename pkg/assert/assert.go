package assert

import (
	"fmt"
	"runtime"
)

func Assert(condition bool, format string, args ...any) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("Assertion failed at %s:%d: %s", file, line, fmt.Sprintf(format, args...)))
	}
}
