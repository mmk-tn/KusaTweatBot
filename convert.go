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
	output = strings.Replace(output, "class=\"mx-auto js-calendar-graph-svg\">",
		"class=\"mx-auto js-calendar-graph-svg\"><style type=\"text/css\">.month, .wday{fill:#767676;font-size:9px;}</style><rect width=\"100%\" height=\"100%\" fill=\"white\" />", 1)
	file.Write(([]byte)(output))
}
