package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func md2html(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.MathJax | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	html_flags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: html_flags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func convert_file(filepath string) error {
	if !strings.HasSuffix(filepath, ".md") {
		return fmt.Errorf("file is not markdown")
	}
	fmt.Printf("Converting %s ", filepath)

	md_file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	filename := md_file.Name()
	if err != nil {
		fmt.Printf("error on opening %s\n%s\n", filepath, err)
	}

	md, err := io.ReadAll(md_file)
	if err != nil {
		fmt.Printf("error on reading %s\n%s\n", filepath, err)
	}

	md_file.Close()
	fmt.Printf(".")

	html := md2html(md)
	html_filename := "out/" + RemoveExt(filename) + ".html"
	os.MkdirAll(path.Dir(html_filename), os.ModePerm)
	fmt.Printf(".")

	html_file, err := os.OpenFile(html_filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("error on opening %s\n%s\n", html_filename, err)
		return err
	}

	_, err = html_file.Write(html)
	if err != nil {
		fmt.Printf("error on writing to %s\n%s\n", html_filename, err)
		return err
	}
	fmt.Printf(".")

	html_file.Close()
	fmt.Printf(" Done\n")
	return nil
}

func convert_dir(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			convert_file(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error on walking path %s\n%s\n", path, err)
	}
}
