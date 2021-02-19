package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	_, _ = r.Body.Read(body)
	_ = json.Unmarshal(body, &post)
	fmt.Println(post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	_, _ = r.Body.Read(body)
	var post Post
	_ = json.Unmarshal(body, &post)
	err = post.create()
	if err != nil {
		return err
	}
	w.WriteHeader(200)
	return
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)

	if err != nil {
		return err
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
