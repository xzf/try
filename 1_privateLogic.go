package try

import (
	"bytes"
	"runtime"
)

func readStack() string {
	content := make([]byte, 102400)
	index := runtime.Stack(content, true)
	if index > 0 {
		end := bytes.Index(content[1:], []byte(`
goroutine`))
		if end > 0 {
			content = content[:end]
		}
	}
	return string(content)
}
