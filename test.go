package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {

	// Get SVG Data
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file.")
	}
	url := os.Getenv("github_url")
	html := getSvgHTML(url)

	// Create SVG Image
	html2svg(html, "Kusa.svg")

	// Convert SVG to PNG
	err2 := exec.Command("cairosvg", "-f", "png", "-o", "Kusa.png", "Kusa.svg").Run()
	if err2 != nil {
		fmt.Println("Error convert svg to png")
	}

	// Tweet Create Image
	tweetWithImage("Tweet from KusaTweetBot", "Kusa.png")
}
