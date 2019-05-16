package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
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

func encode() string {
	file, _ := os.Open("Test.png")
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

func tweatWithImage() {
	loadEnv()

	anaconda.SetConsumerKey(os.Getenv("consumer_key"))
	anaconda.SetConsumerSecret(os.Getenv("consumer_secret"))
	api := anaconda.NewTwitterApi(os.Getenv("access_token"),
		os.Getenv("access_token_secret"))

	image := encode()

	media, _ := api.UploadMedia(image)

	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)
	tweet, err := api.PostTweet("Tweat with a image from anaconda.", v)
	if err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println("Tweat Message:", tweet.Text)
	}
}
