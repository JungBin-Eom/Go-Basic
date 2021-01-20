package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	// 업로드된 파일이 request를 통해 전송
	uploadFile, header, err := r.FormFile("upload_file") // input의 id로 구분
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	dirname := "./uploads"
	os.MkdirAll(dirname, 0777) // 디렉토리 없으면 만들기
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath) // 비어있는 파일 생성
	defer file.Close()               // 연 파일은 defer로 프로그램 종료 전 항상 닫아줄 것
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)
}
