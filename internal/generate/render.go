package generate

import (
	"fmt"
)

func base_template(title string, content string) string {
	return fmt.Sprintf(
		`<html>
	<head>
		<title>%s</title>
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<script>
		MathJax = { tex: { inlineMath: [['\\(', '\\)'], ['\\(', '\\)']] } };
		</script>
		<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-chtml.js">
		</script>
	</head>
	<body>
		%s
	</body>
</html>`,
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

func render_post(post Post) string {
	return base_template(post.Title, post.Content)
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
