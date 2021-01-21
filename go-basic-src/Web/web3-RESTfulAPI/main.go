package main

import (
	"Go-Basic/go-basic-src/Web/web3-RESTfulAPI/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
