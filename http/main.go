package main

import (
	"net/http"
	"time"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(30) * time.Second)
	w.Write([]byte("Hello, world!"))
}

func main() {
	//http.Handle("/", &helloHandler{})
	http.ListenAndServe(":12345", http.FileServer(http.Dir(".")))
}
