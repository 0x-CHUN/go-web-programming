package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	_ = server.ListenAndServe()
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		_, _ = fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	_, _ = fmt.Fprintln(w, c1)
	_, _ = fmt.Fprintln(w, cs)
}

func setCookie(w http.ResponseWriter, _ *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go web programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}
