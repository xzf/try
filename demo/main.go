package main

import (
	"fmt"
	"github.com/xzf/try"
)

func main() {
	try.Try(func() {
		panic("j4l209vi3o")
	}).Catch(func(info interface{}) {
		fmt.Println(info)
	})
}
