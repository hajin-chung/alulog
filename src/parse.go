package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v3"
)

func ParsePostFile(path string) (Post, error) {
	extensions := parser.CommonExtensions | parser.MathJax | parser.NoEmptyLineBeforeBlock
	parser := parser.NewWithExtensions(extensions)
	md_file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("error on opening %s\n%s\n", path, err)
	}

	md, err := io.ReadAll(md_file)
	if err != nil {
		fmt.Printf("error on reading %s\n%s\n", path, err)
	}

	md_file.Close()

	doc := parser.Parse(md)
	post, err := ParseAst(doc)
	if err != nil {
		fmt.Printf("error on parsing ast %s\n%s\n", path, err)
	}
	return post, nil
}

func ParseAst(doc ast.Node) (Post, error) {
	post := Post{}
	post.Content = doc
	children := doc.GetChildren()
	if len(children) == 0 {
		return post, fmt.Errorf("no chidren")
	}

	data_block, ok := children[0].(*ast.CodeBlock)
	if !ok {
		return post, fmt.Errorf("no metadata")
	}
	data_str := data_block.Literal
	err := yaml.Unmarshal([]byte(data_str), &post)
	fmt.Printf("%+v\n", post)
	doc.SetChildren(children[1:])

	return post, err
}

