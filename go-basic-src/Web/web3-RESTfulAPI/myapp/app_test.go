package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal(string(data), "No Users")
}

func TestGetUser89Info(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// GET method 사용
	res, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:89")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"ricky", "last_name":"eom", "email":"ricky@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(res.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// 생성하기 전에 지우기
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:1")

	// 새 User 생성
	res, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"ricky", "last_name":"eom", "email":"ricky@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	// 생성한 User 지우기
	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "Deleted User Id:1")
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// user가 없을 때 update
	req, _ := http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(`{"id":1, "first_name":"updated",	"last_name":"updated",	"email":"updated@gmail.com"}`))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:1")

	// user 생성
	res, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"ricky", "last_name":"eom", "email":"ricky@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	updateStr := fmt.Sprintf(`{"id":%d, "first_name":"bob", "last_name":"torres"}`, user.ID)

	// user가 있을 때 update
	req, _ = http.NewRequest("PUT", ts.URL+"/users", strings.NewReader(updateStr))
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	updatedUser := new(User)
	err = json.NewDecoder(res.Body).Decode(updatedUser)
	assert.NoError(err)
	assert.Equal(updatedUser.ID, user.ID)
	assert.Equal("bob", updatedUser.FirstName)
	assert.Equal("torres", updatedUser.LastName)
	assert.Equal(user.Email, updatedUser.Email)
}

func TestUsersWithUsersData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"ricky", "last_name":"eom", "email":"ricky@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	res, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"venny", "last_name":"kwak", "email":"venny@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	res, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	users := []*User{}
	err = json.NewDecoder(res.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))
}
