package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file.")
	}
}

func tweat() {
	loadEnv()

	anaconda.SetConsumerKey(os.Getenv("consumer_key"))
	anaconda.SetConsumerSecret(os.Getenv("consumer_secret"))
	api := anaconda.NewTwitterApi(os.Getenv("access_token"),
		os.Getenv("access_token_secret"))

	tweet, err := api.PostTweet("Hello World from anaconda with godotenv.", nil)
	if err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println("Tweat Message:", tweet.Text)
	}
}
