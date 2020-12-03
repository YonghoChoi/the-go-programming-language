package main

import (
	"bufio"
	"fmt"
	"os"
)

// 파일에 중복된 행을 카운트하고, 파일이 없는 경우 표준 입력으로 받음
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 : %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 스트리밍 방식으로 파일의 라인을 반복문을 돌 때마다 파일에서 필요한 행을 로드
// 파일 크기가 아무리 커도 실행이 가능
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			break
		}

		counts[line]++
	}
}
