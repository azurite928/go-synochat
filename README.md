# go-synochat

[![Go Report Card](https://goreportcard.com/badge/github.com/azurite928/go-synochat)](https://goreportcard.com/report/github.com/azurite928/go-synochat)

A Go client for Synology Chat's webhooks
API [Synology Chat's WebHooks API](https://kb.synology.com/en-global/DSM/help/Chat/chat_integration).

## Installation

```sh
$ go get -u github.com/azurite928/go-synochat
```

## Usage

Basic text message

```go
webhookURL := "https://DS_IP/webapi/entry.cgi?api=SYNO.Chat.External&XXXXX"
message := synochat.Message{
Text: "Hello Synology Chat!",
}
err := synochat.Send(webhookURL, message)
if err != nil {
fmt.Println(err)
}
```