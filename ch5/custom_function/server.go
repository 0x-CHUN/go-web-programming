package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, _ *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("ch5/custom_function/tmpl.html")
	_ = t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}
