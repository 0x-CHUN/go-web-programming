package main

import (
	"html/template"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}

func process(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("ch5/trigger_template/tmpl.html")
	_ = t.Execute(w, "Hello World!")
}
