package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// URL을 병렬로 반입하고 시간과 크기를 기록
// go run fetchall.go https://google.co.kr https://www.naver.com https://yongho1037.tistory.com
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 고루틴 시작
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // ch 채널에서 수신
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		//ch <- fmt.Sprint(err)
		ch <- err.Error()
		return
	}

	// body의 데이터를 모두 읽지 않으면 메모리 누수가 발생할 수 있기 때문에 아래 로직으로 body를 끝까지 읽고 버림
	// 참고 : https://yongho1037.tistory.com/815, https://medium.com/@xoen/golang-read-from-an-io-readwriter-without-loosing-its-content-2c6911805361
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
