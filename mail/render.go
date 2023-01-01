package mail

type MarkdownEmail interface {
	render(string) string
}

func render(input MarkdownEmail) string {
	return input
}
