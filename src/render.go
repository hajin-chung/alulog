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
			`<li><a href="/post/%s.html">%s - %s</a></li>`, 
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

func render_tags(tag_map map[string][]Post) string {
	tags_list := "<ul>"
	for tag, posts := range tag_map {
		tags_list +=
			fmt.Sprintf(
				`<li><a href="/tag/%s.html">%s (%d)</a></li>`,
				tag, tag, len(posts),
			)
	}
	return base_template("Alulog - tags", tags_list)
}

func render_tag(tag string, posts []Post) string {
	post_list := "<ul>"
	for _, post := range posts {
		post_list += 
			fmt.Sprintf(
			`<li><a href="/post/%s.html">%s - %s</a></li>`, 
			SanitizeTitle(post.Title), post.Created, post.Title,
		)
	}
	post_list += "</ul>"
	return base_template(fmt.Sprintf("Alulog - %s", tag), post_list)
}
