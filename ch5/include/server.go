package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("ch5/include/t1.html", "ch5/include/t2.html")
	_ = t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}
