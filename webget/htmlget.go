package webget

import (
	"fmt"
	"github.com/go-chat-bot/bot"
    "strings"
    "net/url"
    "github.com/PuerkitoBio/goquery"
)

    //func gif(command *bot.Cmd) (msg string, err error) {
func htmlGet(cmd *bot.Cmd) (msg string, err error) {

    if len(cmd.Args) != 2 {
        msg = "Invalid number of arguments"
        return
    }

    pageUrl, err := url.Parse( cmd.Args[0])
    if err != nil { return }

    selector := cmd.Args[1]

    doc, err := goquery.NewDocument( pageUrl.String())
    if err != nil { return }

    firstItem := doc.Find( selector).First()

    msg = strings.TrimSpace( firstItem.Text())

    return

    fmt.Println( pageUrl)
    fmt.Println( selector)

    return
}

func init() {
	bot.RegisterCommand(
		"htmlget",
		"Get the value of a specific HTML element",
		"",
		htmlGet)
}
