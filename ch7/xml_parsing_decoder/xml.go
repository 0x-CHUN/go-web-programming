package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("ch7/xml_parsing_decoder/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoer := xml.NewDecoder(xmlFile)
	for {
		t, err := decoer.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				_ = decoer.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}
}
