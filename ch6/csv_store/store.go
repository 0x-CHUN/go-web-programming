package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("./ch6/csv_store/posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{ID: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{ID: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("./ch6/csv_store/posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{ID: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].ID)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
