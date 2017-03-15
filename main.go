package main

import (
	"encoding/xml"
	"fmt"
	// "os"
	"io/ioutil"
	"regexp"
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
	URL          string `xml:"url"`
	Id           int    `xml:"id"`
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

	var contentFile ContentFile
	for _, value := range blogposts.Posts {

		contentFile.Author = value.Author
		contentFile.Date = value.Created.Date
		contentFile.DateUpdate = value.Updated.Date
		contentFile.Title = value.Title
		contentFile.Filename = strings.ToLower(value.URL)

		// de blogs
		searchTermDe := `http://www.mynewsdesk.com/de/nimirum/blog_posts/`
		reDe := regexp.MustCompile(searchTermDe)
		reDeSlice := reDe.FindStringSubmatch(string(contentFile.Filename))

		if len(reDeSlice) > 0 {
			fWithoutUrlDe := strings.TrimPrefix(contentFile.Filename, reDeSlice[0])
			contentFile.Filename = fWithoutUrlDe[:len(fWithoutUrlDe)-6]
		}

		// uk blogs
		searchTermUk := `http://www.mynewsdesk.com/uk/nimirum/blog_posts/`
		reUk := regexp.MustCompile(searchTermUk)
		reUkSlice := reUk.FindStringSubmatch(string(contentFile.Filename))

		if len(reUkSlice) > 0 {
			fWithoutUrlUk := strings.TrimPrefix(contentFile.Filename, reUkSlice[0])
			contentFile.Filename = fWithoutUrlUk[:len(fWithoutUrlUk)-6]

		}

		fmt.Println(contentFile.Filename)

	}

	return blogposts, nil

}

func makeContentFile(BlogPosts) (ContentFile, error) {

	return ContentFile{}, nil
}
