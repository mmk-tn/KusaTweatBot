package main

import (
	"fmt"
	"os"
	"strings"
)

func html2svg(html string, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	output := strings.Replace(html, "<svg", "<svg xmlns=\"http://www.w3.org/2000/svg\"", 1)
	file.Write(([]byte)(output))
}
