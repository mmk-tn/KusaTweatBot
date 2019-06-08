package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Set Logging
	stdout, stderr, err := initOut("./log/", "KusaTweetBot")
	if err != nil {
		fmt.Fprintf(stderr, "Can't creata file: %v\n", err)
	}
	fmt.Fprintf(stdout, "%s - Start KusaTweetBot.\n", time.Now())

	// Get SVG Data
	err = godotenv.Load()
	if err != nil {
		fmt.Fprintln(stderr, "Error loading env file.")
	}
	url := os.Getenv("github_url")
	html := getSvgHTML(url)

	// Create SVG Image
	html2svg(html, "Kusa.svg")

	// Convert SVG to PNG
	err = exec.Command("cairosvg", "-f", "png", "-o", "Kusa.png", "Kusa.svg").Run()
	if err != nil {
		fmt.Fprintln(stderr, "Error convert svg to png.")
	}

	// Tweet Create Image
	message := fmt.Sprint("[", time.Now().Format("2006-01-02"), "]", "by KusaTweetBot\n", url)
	tweetWithImage(message, "Kusa.png")

	fmt.Fprintf(stdout, "%s - Finish KusaTweetBot.\n", time.Now())
}
