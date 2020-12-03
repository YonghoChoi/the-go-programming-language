package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	// Scanner는 줄 단위로 들어오는 입력을 처리하는 가장 쉬운 방법
	// 표준 입력을 받음
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			break
		}

		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
