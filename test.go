package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// tweet()
	// tweetWithImage()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file.")
	}
	url := os.Getenv("github_url")

	fmt.Println("Get Data:", getSvgHTML(url))
}
