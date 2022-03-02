package utility

import (
	"fmt"
	"strings"
)

func CreateMarkdownList(items []string) string {
	list := ""
	for _, item := range items {
		list += fmt.Sprintf("-\t%s\n", item)
	}

	return list
}

func CreateMarkdownHeading(text string, level int) string {
	return fmt.Sprintf("%s %s", strings.Repeat("#", level), text)
}

func CreateMarkdownHeadings(text []string, level int) string {
	headings := ""
	for _, t := range text {
		headings += fmt.Sprintf("%s\n", CreateMarkdownHeading(t, level))
	}
	return headings
}
