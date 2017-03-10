package main

import (
	"encoding/xml"
	"fmt"
	// "os"
	"io/ioutil"
)

func main() {

	query()

}

type BlogPost struct {
	TotalCount string   `xml:"total-count,attr"`
	SourceId   []string `xml:"item>source_id"`
}

func query() (BlogPost, error) {

	// file, err := os.Open("posts/blogposts.xml")
	file, err := ioutil.ReadFile("posts/blogposts.xml")

	if err != nil {
		return BlogPost{}, err
	}

	var post BlogPost
	if err := xml.Unmarshal([]byte(file), &post); err != nil {
		fmt.Println(BlogPost{}, err)
		return BlogPost{}, err

	}

	fmt.Println(len(post.SourceId))
	return post, nil

}
