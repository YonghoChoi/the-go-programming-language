package main

import (
	"fmt"
	"log"
	"net/http"
)

// go run server1.go
func main() {
	http.HandleFunc("/", handler) // 각 요청은 핸들러를 호출
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
