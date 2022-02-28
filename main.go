package main

import (
	"os"

	"github.com/Tch1b0/readcli/cmd"
)

func main() {
	// cmd.Execute()
	s, _ := cmd.CreateReadme().Render()
	os.WriteFile("README.md", s, 0600)
}
