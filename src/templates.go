package main

import "fmt"

func BaseTemplate(title string, content string) string {
	return fmt.Sprintf(`<html>
  <head>
    <title>%s</title>
  </head>
  <body>
    %s
  </body>
</html>`, title, content)
}

func PostTemplate(title string, content string) string {
	return BaseTemplate(title, content)
}
