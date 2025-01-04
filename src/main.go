package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

type Post struct {
	Title string `yaml:"title"`
	Created string `yaml:"created"`
	Updated string `yaml:"updated"`
	Tags []string `yaml:"tags"`
	Content ast.Node
}

func main() {
	html_flags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: html_flags}
	renderer := html.NewRenderer(opts)

	tag_map := map[string][]Post{}
	posts := []Post{}

	entries, err := os.ReadDir("posts")
	if err != nil {
		fmt.Printf("error on reading dir ./posts\n%s\n", err)
		return
	}

	for _, entry := range entries {
		path := "posts/" + entry.Name();
		fmt.Printf("%s\n", path)
		if entry.IsDir() {
			continue
		}
		post, err := ParsePostFile(path)
		if err != nil {
			fmt.Printf("error on parsing %s\n%s\n", entry.Name(), err)
		}
		posts = append(posts, post)
	}

	sort.Slice(posts,	func(i, j int) bool {
		it, _ := time.Parse("2006-01-02", posts[i].Created)
		jt, _ := time.Parse("2006-01-02", posts[j].Created)
		return it.After(jt)
	})

	for _, post := range posts {
		for _, tag := range post.Tags {
			tag_map[tag] = append(tag_map[tag], post)
		}
	} 

	os.Mkdir("out/", os.ModePerm)
	os.Mkdir("out/post", os.ModePerm)
	os.Mkdir("out/tag", os.ModePerm)

	index_html := render_index(posts)
	os.WriteFile("out/index.html", []byte(index_html), os.ModePerm)

	tags_html := render_tags(tag_map)
	os.WriteFile("out/tags.html", []byte(tags_html), os.ModePerm)

	for tag, posts := range tag_map {
		tag_html := render_tag(tag, posts)
		os.WriteFile(
			fmt.Sprintf("out/tag/%s.html", tag), 
			[]byte(tag_html), 
			os.ModePerm,
		)
	}

	for _, post := range posts {
		post_html := render_post(post, renderer)
		out_path := SanitizeTitle(fmt.Sprintf("out/post/%s.html", post.Title))
		os.WriteFile(out_path, []byte(post_html), os.ModePerm)
	}
}
