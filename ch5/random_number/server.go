package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}

func process(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("ch5/random_number/tmpl.html")
	rand.Seed(time.Now().Unix())
	_ = t.Execute(w, rand.Intn(10) > 5)
}
