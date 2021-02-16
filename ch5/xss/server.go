package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	t, _ := template.ParseFiles("ch5/xss/tmpl.html")
	_ = t.Execute(w, r.FormValue("comment"))
	_ = t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("ch5/xss/form.html")
	_ = t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/", form)
	_ = server.ListenAndServe()
}
