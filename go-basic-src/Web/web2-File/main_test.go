package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "~/ricky/Desktop/캐리커쳐.PNG"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf) // 버퍼 전달해서 이 버퍼에 데이터 저장
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))

	assert.NoError(err)

	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())
	uploadsHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath) // file 정보 가져옴

	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}
