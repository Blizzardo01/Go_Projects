package main

import ("fmt"
        "time"
        "github.com/slack-go/slack"
      )

func main() {

  authtoken := "xoxb-2993577612355-2993554338050-qcMcuShIClrso9QlCZBZ6Q4K"
  channeltoken := "C02V7H0GP8B"

  client := slack.New(authtoken, slack.OptionDebug(true))

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
    channeltoken,
    slack.MsgOptionAttachments(attachment),
    )

    if err != nil {
      panic(err)
    }

  fmt.Printf("Message sent at %s", timestamp)

}
