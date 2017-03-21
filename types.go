package main

type Frontmatter struct {
	Title        string
	Id           string
	Author       string
	Date         string
	DateUpdate   string
	Language     string
	Summary      string
	Content      string
	Image        string
	ImageCaption string
	Tags         []string
}

type ContentFile struct {
	Frontmatter
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

type Pressrelease struct {
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

type Pressreleases struct {
	TotalCount string         `xml:"total-count,attr"`
	Releases   []Pressrelease `xml:"item"`
}

type News struct {
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

type Pressnews struct {
	TotalCount string `xml:"total-count,attr"`
	PressNews  []News `xml:"item"`
}
