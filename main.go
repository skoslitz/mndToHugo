package main

import (
	"encoding/xml"
	"fmt"
	//"io"
	"io/ioutil"
	//"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"text/template"
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

	postQuantities := len(blogposts.Posts)
	var contentFile ContentFile
	for index, value := range blogposts.Posts {

		quote := regexp.MustCompile(`"`)
		table := regexp.MustCompile(`(?s)(<table>.*<\/table>)`)
		newLine := regexp.MustCompile(`(?s)(\n|\r|\r\n)`)

		contentFile.Title = strings.TrimSpace(quote.ReplaceAllString(value.Title,
			"\\\""))
		contentFile.Id = "B_" + strconv.Itoa(postQuantities-index)
		contentFile.Author = value.Author
		contentFile.Date = value.Created.Date
		contentFile.DateUpdate = value.Updated.Date
		contentFile.Language = value.Language
		contentFile.Summary = strings.TrimSpace(quote.ReplaceAllString(value.Summary,
			"\\\""))
		contentFile.Summary = newLine.ReplaceAllString(contentFile.Summary,
			" ")
		contentFile.Image = strings.TrimSpace(value.Image)
		contentFile.ImageCaption = strings.TrimSpace(quote.ReplaceAllString(value.ImageCaption,
			"\\\""))
		contentFile.Tags = value.Tags

		contentFile.Content = table.ReplaceAllString(value.Body, "")
		contentFile.Content = strings.TrimSpace(quote.ReplaceAllString(contentFile.Content,
			"\\\""))

		/* build filenames*/
		contentFile.Filename = strings.ToLower(value.URL)

		/* set img to /static path */
		if len(contentFile.Image) > 5 {
			fileLength := len(contentFile.Image)
			fileName := contentFile.Id
			fileExt := contentFile.Image[fileLength-3 : fileLength]
			filePath := path.Join(contentFile.Id, fileName+"."+fileExt)

			contentFile.Image = filePath
		}

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

		/* build .md from Contentfile */

		blogpost := Frontmatter{
			Title:        contentFile.Title,
			Id:           contentFile.Id,
			Author:       contentFile.Author,
			Date:         contentFile.Date,
			DateUpdate:   contentFile.DateUpdate,
			Language:     contentFile.Language,
			Summary:      contentFile.Summary,
			Content:      contentFile.Content,
			Image:        contentFile.Image,
			ImageCaption: contentFile.ImageCaption,
			Tags:         contentFile.Tags,
		}
		tmpl, err := template.ParseFiles("blogpost.md")

		if err != nil {
			panic(err)
		}

		mdFileName := contentFile.Id + "-" + contentFile.Filename + ".md"
		mdFilePath := path.Join("content/blog/", mdFileName)

		mdFile, err := os.Create(mdFilePath)
		if err != nil {
			panic(err)
		}

		err = tmpl.ExecuteTemplate(mdFile, "blogpost.md", blogpost)
		if err != nil {
			panic(err)
		}

		defer mdFile.Close()

		/* prepare post assets and copy them*/
		/*staticDir := path.Join("static/img/blog/", contentFile.Id)
		os.MkdirAll(staticDir, 0755)

		if len(contentFile.Image) > 5 {

			resp, err := http.Get(contentFile.Image)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			fileLength := len(contentFile.Image)
			fileName := contentFile.Id
			fileExt := contentFile.Image[fileLength-3 : fileLength]
			filePath := path.Join("static/img/blog/", contentFile.Id, fileName+"."+fileExt)

			imgfile, err := os.Create(filePath)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(imgfile, resp.Body)
			if err != nil {
				panic(err)
			}

			defer imgfile.Close()
			fmt.Println("\n Img copy done")
		}*/

	}

	return blogposts, nil

}

func makeContentFile(BlogPosts) (ContentFile, error) {

	return ContentFile{}, nil
}
