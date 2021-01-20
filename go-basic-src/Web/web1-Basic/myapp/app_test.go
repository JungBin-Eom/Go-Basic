// go는 파일이름 뒤에 '_test'를 붙이면 test code가 된다.

package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) { // 함수 이름앞에 'Test', 인자로 '*testing.T' 전달함으로써 테스트 수행
	assert := assert.New(t)

	// httptest 패키지의 NewRecorder/NewRequest로 가상의 response와 request를 만듦
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req) // mux를 만들어서 호출해야 target에 맞추어 호출

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body) // res.Body는 버퍼형태이므로 ioutil 필요
	assert.Equal("Hello World", string(data))
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed!", res.Code)
	// }
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=ricky", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello ricky", string(data))
}

func TestFooPathHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooPathHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"ricky", "last_name":"eom", "email":"ricky@gmail"}`))
	// string을 io.Reader로 바꾸어 전달해야 함(Buffer)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)

	assert.Nil(err)
	assert.Equal("ricky", user.FirstName)
	assert.Equal("eom", user.LastName)
}
