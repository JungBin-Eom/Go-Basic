package main

import (
	"Go-Basic/go-basic-src/Web/web1/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler()) // 웹서버 실행, 지정한 port에서 request 대기
}
