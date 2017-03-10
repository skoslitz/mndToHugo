package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	query()

}

/*type Item []struct {
	Language string `json: "language"`
}*/

/*
type Items struct {
	Item       []interface{} `json: "item"`
	TotalCount string        `json: total_count"`
}*/

type BlogPost struct {
	Items struct {
		Item []struct {
			Published string `json: "published_at"`
			Language  string `json: "language"`
		} `json: "item"`
		TotalCount string `json: total_count"`
	} `json: "items"`
}

func query() (BlogPost, error) {

	file, err := os.Open("posts/blog_posts.json")
	if err != nil {
		return BlogPost{}, err
	}

	defer file.Close()

	var d BlogPost
	if err := json.NewDecoder(file).Decode(&d); err != nil {
		fmt.Println(BlogPost{}, err)
		return BlogPost{}, err

	}

	fmt.Println(d.Items.TotalCount)
	return d, nil

}
