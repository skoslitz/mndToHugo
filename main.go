package main

import (
	"encoding/xml"
	"fmt"
	// "os"
	"io/ioutil"
	"strings"
)

func main() {

	parseBlog()

}

type Frontmatter struct {
	Title        string
	Id           string
	Author       string
	Date         string
	DateUpdate   string
	Language     string
	Summary      string
	Image        string
	ImageCaption string
	Tags         []string
}

type ContentFile struct {
	Frontmatter
	Content  string
	Filename string
}

type Datetime struct {
	Date string `xml:"datetime,attr"`
}

type Post struct {
	Title        string `xml:"header"`
	Author       string
	Created      Datetime `xml:"created_at"`
	Updated      Datetime `xml:"updated_at"`
	Language     string   `xml:"language"`
	Summary      string   `xml:"summary"`
	Body         string   `xml:"body"`
	Image        string   `xml:"image"`
	ImageCaption string   `xml:"image_caption"`
	Tags         []string `xml:"tags>tag"`
}

type BlogPosts struct {
	TotalCount string `xml:"total-count,attr"`
	Posts      []Post `xml:"item"`
}

func parseBlog() (BlogPosts, error) {

	file, err := ioutil.ReadFile("posts/blogposts.xml")

	if err != nil {
		return BlogPosts{}, err
	}

	var blogposts BlogPosts
	if err := xml.Unmarshal([]byte(file), &blogposts); err != nil {
		fmt.Println(BlogPosts{}, err)
		return BlogPosts{}, err

	}

	fmt.Println(blogposts.Posts[58].Title)
	fmt.Println(blogposts.Posts[59].Title)

	var contentFile ContentFile
	for _, value := range blogposts.Posts {

		contentFile.Author = value.Author
		contentFile.Date = value.Created.Date
		contentFile.DateUpdate = value.Updated.Date
		contentFile.Filename = strings.ToLower(value.Title)
		fmt.Println(value.Language, contentFile.Filename)
	}

	return blogposts, nil

}

func makeContentFile(BlogPosts) (ContentFile, error) {

	return ContentFile{}, nil
}
