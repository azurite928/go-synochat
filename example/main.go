package main

import (
	"fmt"

	"github.com/azurite928/go-synochat"
)

func main() {
	webhookURL := "https://DS_IP/webapi/entry.cgi?api=SYNO.Chat.External&XXXXX"
	message := synochat.Message{
		Text: "Hello Synology Chat!",
	}
	err := synochat.Send(webhookURL, message)
	if err != nil {
		fmt.Println(err)
	}
}
