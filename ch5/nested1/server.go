package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("ch5/nested1/layout.html")
	_ = t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}
