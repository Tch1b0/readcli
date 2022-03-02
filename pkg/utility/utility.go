package utility

import "fmt"

func CreateMarkdownList(items []string) string {
	list := ""
	for _, item := range items {
		list += fmt.Sprintf("-\t%s\n", item)
	}

	return list
}
