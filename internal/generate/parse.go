package generate

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	gparser "github.com/yuin/goldmark/parser"
)

func ParsePostFile(path string, parser goldmark.Markdown) (Post, error) {
	md_file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("error on opening %s\n%s\n", path, err)
	}

	md, err := io.ReadAll(md_file)
	if err != nil {
		fmt.Printf("error on reading %s\n%s\n", path, err)
	}

	md_file.Close()

	var buf bytes.Buffer
	context := gparser.NewContext()
	err = parser.Convert(md, &buf, gparser.WithContext(context))
	if err != nil {
		fmt.Printf("error on converting %s\n%s\n", path, err)
	}

	metadata := meta.Get(context)
	fmt.Printf("%+v\n", metadata)
	title, ok := metadata["title"].(string)
	if !ok {
		title = ""
	}
	created, ok := metadata["created"].(string)
	if !ok {
		created = ""
	}
	updated, ok := metadata["updated"].(string)
	if !ok {
		updated = ""
	}
	tags_interface := metadata["tags"].([]interface{})
	tags := make([]string, len(tags_interface))
	for i, v := range tags_interface {
		tags[i] = v.(string)
	}
	post := Post{
		Title:   title,
		Created: created,
		Updated: updated,
		Tags:    tags,
		Content: buf.String(),
	}
	return post, nil
}
