package main

import (
	"fmt"
	"os"
	"strings"
)

// command line 인수 출력
// go run echo.go arg1 arg2
func main() {
	var s, sep string
	//s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
