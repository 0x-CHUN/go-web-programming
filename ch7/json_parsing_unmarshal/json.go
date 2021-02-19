package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

func main() {
	jsonFile, err := os.Open("ch7/json_parsing_unmarshal/post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	var post Post
	_ = json.Unmarshal(jsonData, &post)
	fmt.Println(post)
	fmt.Println(post.Id)
	fmt.Println(post.Content)
	fmt.Println(post.Author.Id)
	fmt.Println(post.Author.Name)
	fmt.Println(post.Comments)
	fmt.Println(post.Comments[0].Id)
	fmt.Println(post.Comments[0].Content)
	fmt.Println(post.Comments[0].Author)
}
