package parser

import (
	"html-link-parser/internal/parser/fileType"
	"html-link-parser/internal/parser/htmlFile"
	"html-link-parser/internal/parser/link"
)

func HrefParse(path string, t fileType.FileType) []link.Link {
	file := htmlFile.NewHtmlFile(path, t)
	return file.GetLinks()
}
