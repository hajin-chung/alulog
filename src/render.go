package main

import (
	"fmt"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func base_template(title string, content string) string {
	return fmt.Sprintf(
		`<html><head><title>%s</title></head><body>%s</body></html>`, 
		title, content,
	)
}

func render_index(posts []Post) string {
	post_list := "<ul>"
	for _, post := range posts {
		post_list += 
			fmt.Sprintf(
			`<li><a href="/post/%s">%s - %s</a></li>`, 
			SanitizeTitle(post.Title), post.Created, post.Title,
		)
	}
	post_list += "</ul>"
	return base_template("Alulog", post_list)
}

func render_post(post Post, renderer *html.Renderer) string {
	content := string(markdown.Render(post.Content, renderer))
	return base_template(post.Title, content)
}
