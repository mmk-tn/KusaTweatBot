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

func encode(imPath string) string {
	file, _ := os.Open(imPath)
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

func tweetWithImage(message string, imPath string) {
	loadEnv()

	anaconda.SetConsumerKey(os.Getenv("consumer_key"))
	anaconda.SetConsumerSecret(os.Getenv("consumer_secret"))
	api := anaconda.NewTwitterApi(os.Getenv("access_token"),
		os.Getenv("access_token_secret"))

	image := encode(imPath)

	media, _ := api.UploadMedia(image)

	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)
	tweet, err := api.PostTweet(message, v)
	if err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println("Tweet Message:", tweet.Text)
	}
}
