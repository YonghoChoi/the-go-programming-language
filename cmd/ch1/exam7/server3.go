package main

import (
	"fmt"
	"log"
	"net/http"
)

// go run server2.go
// 핸들러 Request의 상세 정보 확인
func main() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	for k, v := range r.Header {
		// %q는 문자열에 따옴표를 포함해서 표시
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// URL에서 raw 쿼리를 구문 분석하고 r.Form에 데이터를 갱신
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
