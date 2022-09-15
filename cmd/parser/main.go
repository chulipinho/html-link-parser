package main

import (
	"flag"
	"fmt"
	"html-link-parser/internal/parser"
	"html-link-parser/internal/parser/link"
	"html-link-parser/internal/pkg/errors"
)

func main() {
	filePath := flag.String("f", "", "Path to HTML file to be read")
	linkPath := flag.String("p", "", "Link to HTML page to be read")
	flag.Parse()

	if *filePath != "" && *linkPath != "" {
		errors.ThrowError("Please use only one flag")
	} else if *filePath == "" && *linkPath == "" {
		errors.ThrowError("Please provide a file or link")
	}

	var links []link.Link
	switch {
	case *filePath != "":
		links = parser.HrefParse(*filePath, 1)
	case *linkPath != "":
		links = parser.HrefParse(*linkPath, 0)
	}

	printLinks(links)
}

func printLinks(links []link.Link) {
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
