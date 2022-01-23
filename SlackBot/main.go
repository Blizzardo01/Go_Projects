package main

import ("fmt"
        "os"
        "time"

        "github.com/joho/godotenv"
        "github.com/slack-go/slack"
      )

func main() {

  godotenv.Load(".env")

  token := os.Getenv("SLACK_AUTH_TOKEN")
  channel1ID := os.Getenv("SLACK_CHANNEL_ID")

  client := slack.New(token, slack.OptionDebug(true))

  attachment := slack.Attachment{
    Pretext: "Testing Message",
    Text: "blah blah",
    Color: "#36a64f",

    Fields: []slack.AttachmentField{
      {
        Title: "Date",
        Value: time.Now().Format("01-02-2006 15:04:05"),
      },
    },
  }
    _, timestamp, err := client.PostMessage(
    channel1ID,
    slack.MsgOptionAttachments(attachment),
    )

    if err != nil {
      panic(err)
    }

  fmt.Printf("Message sent at %s", timestamp)

}
