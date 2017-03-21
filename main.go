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
	//parsePressreleases()
	//parsePressnews()

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

func parsePressreleases() (Pressreleases, error) {

	file, err := ioutil.ReadFile("posts/pressreleases.xml")

	if err != nil {
		return Pressreleases{}, err
	}

	var pressreleases Pressreleases
	if err := xml.Unmarshal([]byte(file), &pressreleases); err != nil {
		fmt.Println(Pressreleases{}, err)
		return Pressreleases{}, err

	}

	releaseQuantities := len(pressreleases.Releases)
	var contentFile ContentFile
	for index, value := range pressreleases.Releases {

		quote := regexp.MustCompile(`"`)
		newLine := regexp.MustCompile(`(?s)(\n|\r|\r\n)`)

		contentFile.Title = strings.TrimSpace(quote.ReplaceAllString(value.Title,
			"\\\""))
		contentFile.Id = "P_" + strconv.Itoa(releaseQuantities-index)
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
		contentFile.Content = newLine.ReplaceAllString(value.Body,
			" ")
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

		// de releases
		searchTermDe := `http://www.mynewsdesk.com/de/nimirum/pressreleases/`
		reDe := regexp.MustCompile(searchTermDe)
		reDeSlice := reDe.FindStringSubmatch(string(contentFile.Filename))

		if len(reDeSlice) > 0 {
			fWithoutUrlDe := strings.TrimPrefix(contentFile.Filename, reDeSlice[0])
			contentFile.Filename = fWithoutUrlDe[:len(fWithoutUrlDe)-7]
		}

		// uk releases
		searchTermUk := `http://www.mynewsdesk.com/uk/nimirum/pressreleases/`
		reUk := regexp.MustCompile(searchTermUk)
		reUkSlice := reUk.FindStringSubmatch(string(contentFile.Filename))

		if len(reUkSlice) > 0 {
			fWithoutUrlUk := strings.TrimPrefix(contentFile.Filename, reUkSlice[0])
			contentFile.Filename = fWithoutUrlUk[:len(fWithoutUrlUk)-7]

		}

		/* build .md from Contentfile */

		pressrelease := Frontmatter{
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
		tmpl, err := template.ParseFiles("pressrelease.md")

		if err != nil {
			panic(err)
		}

		mdFileName := contentFile.Id + "-" + contentFile.Filename + ".md"
		mdFilePath := path.Join("content/presse/", mdFileName)

		mdFile, err := os.Create(mdFilePath)
		if err != nil {
			panic(err)
		}

		err = tmpl.ExecuteTemplate(mdFile, "pressrelease.md", pressrelease)
		if err != nil {
			panic(err)
		}

		defer mdFile.Close()

		/* prepare post assets and copy them*/
		/*staticDir := path.Join("static/img/presse/", contentFile.Id)
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
			filePath := path.Join("static/img/presse/", contentFile.Id, fileName+"."+fileExt)

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

	return pressreleases, nil
}

func parsePressnews() (Pressnews, error) {
	file, err := ioutil.ReadFile("posts/pressnews.xml")

	if err != nil {
		return Pressnews{}, err
	}

	var pressnews Pressnews
	if err := xml.Unmarshal([]byte(file), &pressnews); err != nil {
		fmt.Println(Pressnews{}, err)
		return Pressnews{}, err

	}

	releaseQuantities := len(pressnews.PressNews)
	var contentFile ContentFile
	for index, value := range pressnews.PressNews {

		quote := regexp.MustCompile(`"`)
		table := regexp.MustCompile(`(?s)(<table>.*<\/table>)`)
		newLine := regexp.MustCompile(`(?s)(\n|\r|\r\n)`)

		contentFile.Title = strings.TrimSpace(quote.ReplaceAllString(value.Title,
			"\\\""))
		contentFile.Id = "N_" + strconv.Itoa(releaseQuantities-index)
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

		// de releases
		searchTermDe := `http://www.mynewsdesk.com/de/nimirum/news/`
		reDe := regexp.MustCompile(searchTermDe)
		reDeSlice := reDe.FindStringSubmatch(string(contentFile.Filename))

		if len(reDeSlice) > 0 {
			fWithoutUrlDe := strings.TrimPrefix(contentFile.Filename, reDeSlice[0])
			contentFile.Filename = fWithoutUrlDe[:len(fWithoutUrlDe)-6]
		}

		// uk releases
		searchTermUk := `http://www.mynewsdesk.com/uk/nimirum/news/`
		reUk := regexp.MustCompile(searchTermUk)
		reUkSlice := reUk.FindStringSubmatch(string(contentFile.Filename))

		if len(reUkSlice) > 0 {
			fWithoutUrlUk := strings.TrimPrefix(contentFile.Filename, reUkSlice[0])
			contentFile.Filename = fWithoutUrlUk[:len(fWithoutUrlUk)-6]

		}

		/* build .md from Contentfile */

		news := Frontmatter{
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
		tmpl, err := template.ParseFiles("pressnews.md")

		if err != nil {
			panic(err)
		}

		mdFileName := contentFile.Id + "-" + contentFile.Filename + ".md"
		mdFilePath := path.Join("content/presse/", mdFileName)

		mdFile, err := os.Create(mdFilePath)
		if err != nil {
			panic(err)
		}

		err = tmpl.ExecuteTemplate(mdFile, "pressnews.md", news)
		if err != nil {
			panic(err)
		}

		defer mdFile.Close()

		/* prepare post assets and copy them*/
		/*staticDir := path.Join("static/img/presse/", contentFile.Id)
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
			filePath := path.Join("static/img/presse/", contentFile.Id, fileName+"."+fileExt)

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

	return pressnews, nil
}
