package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "3",
			Name: "xxx",
		},
	}

	xmlFile, err := os.Create("ch7/xml_creating_encoder/post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
