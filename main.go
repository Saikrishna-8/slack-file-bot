package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6038982846305-6026310033619-97VUMpeYtsBDEXhjatk3xcmD")
	os.Setenv("CHANNEL_ID", "C061F9ZN5KJ")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}

	fileArr := []string{"lets_go_by_alex_edwards_2ed.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n ", file.Name, file.URL)

	}
}
