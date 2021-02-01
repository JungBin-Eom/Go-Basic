package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "ricky", Email: "ricky@gmail.com"}

	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
	rd.JSON(w, http.StatusOK, user)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "ricky", Email: "ricky@gmail.com"}
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	// w.WriteHeader(http.StatusBadRequest)
	// 	// fmt.Fprint(w, err)
	// 	rd.Text(w, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "Ricky")

	// render는 템플릿 등록시 기본 .tmpl 확장자
	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	rd = render.New(render.Options{
		Directory:  "template",                 // Directory로 폴더 추가 등록
		Extensions: []string{".html", ".tmpl"}, // Extensions으로 확장자 추가 등록
		Layout:     "hello",                    // Layout으로 레이아웃 지정
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic() // 기본 파일 서버 기능 제공
	// 로그 기능 제공

	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
