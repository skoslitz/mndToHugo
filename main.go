package main

import (
	"encoding/xml"
	"fmt"
	// "os"
	"io/ioutil"
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

	fmt.Println(len(blogposts.Posts[0].Tags))
	fmt.Println("---")
	fmt.Println(blogposts.Posts[0].Tags)

	return blogposts, nil

}
