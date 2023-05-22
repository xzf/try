# try
try catch for golang 

# demo 
```
package main

import (
	"fmt"
	
	"github.com/xzf/try"
)

func main() {
	try.Try(func() {
		panic("????")
	}).Catch(func(panicInfo interface{}) {
		fmt.Println("catch err", panicInfo)
	})
}
```