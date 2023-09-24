package errcraft

import (
	"bytes"
	"fmt"
	"runtime"
)

// CaptureStack returns a formatted string representation of the call stack.
// The recording starts from the 'start' level and goes up to the 'end' level.
// If 'end' is <= 0, it captures until the end of the stack.
func CaptureStack(start, end int) string {
	var stack bytes.Buffer

	// Iterate through the call stack.
	for depth := start; ; depth++ {
		// If 'end' is set and depth has reached it, break.
		if end > 0 && depth >= end {
			break
		}

		// Retrieve caller information from the current stack depth.
		pc, file, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}

		// Format and write the information to the buffer.
		funcName := runtime.FuncForPC(pc).Name()
		stack.WriteString(fmt.Sprintf("%s:%d %s\n", file, line, funcName))
	}

	return stack.String()
}
