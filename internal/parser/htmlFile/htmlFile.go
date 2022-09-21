package htmlFile

import (
	"html-link-parser/internal/parser/fileType"
	"html-link-parser/internal/parser/link"
	"html-link-parser/internal/pkg/errors"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func NewHtmlFile(path string, t fileType.FileType) *htmlFile {
	f := &htmlFile{
		Path: path,
		Type: t,
	}

	f.getHtmlFile()
	defer f.file.Close()

	n, err := html.Parse(f.file)
	if err != nil {
		errors.ThrowError("Error parsing content no HTML Node. Message: ", err.Error())
	}
	f.Node = n

	return f
}

type htmlFile struct {
	Path string
	Type fileType.FileType
	Node *html.Node
	file io.ReadCloser
}

func (f *htmlFile) getHtmlFile() {
	switch f.Type {
	case fileType.Webpage:
		res, err := http.Get(f.Path)
		if err != nil {
			errors.ThrowError("Error opening webpage. Message: ", err.Error())
		}
		f.file = res.Body
	case fileType.Local:
		file, err := os.Open(f.Path)
		if err != nil {
			errors.ThrowError("Error reading file. Message: ", err.Error())
		}
		f.file = file
	default:
		errors.ThrowError("Something went wrong getting your file.")
	}
}

func (f *htmlFile) GetLinks() []link.Link {
	var links []link.Link

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			link := getAElement(n)
			links = append(links, link)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(f.Node)

	return links
}

func getAElement(node *html.Node) link.Link {
	text := ""

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.TextNode {
			text += n.Data + " "
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(node)

	return parseNode(node, text)
}

func parseNode(hrefNode *html.Node, text string) link.Link {
	var href string

	for _, attr := range hrefNode.Attr {
		if attr.Key == "href" {
			href = attr.Val
			break
		}
	}

	a := link.Link{
		Href: href,
		Text: strings.TrimSpace(text),
	}

	return a
}
