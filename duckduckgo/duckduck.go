package duckduckgo

import (
	"fmt"
	"github.com/go-chat-bot/bot"
    "strings"
    "net/url"
    "github.com/PuerkitoBio/goquery"
)

func ddgSearch(cmd *bot.Cmd) (msg string, err error) {

    searchUrl, _ := url.Parse("https://duckduckgo.com/html/")

    params := url.Values{}
    params.Set("q", cmd.RawArgs)
    searchUrl.RawQuery = params.Encode()

    doc, err := goquery.NewDocument(searchUrl.String())

    if err != nil {
        return
    }

    firstLink := doc.Find("div.results h2 > a").First()

    href := "http://" + strings.TrimSpace(doc.Find("a.result__url").First().Text())

    msg = fmt.Sprintf( "%s: %s", firstLink.Text(), href)

    return
}

func init() {
	bot.RegisterCommand(
		"ddg",
		"Search DuckDuckGo",
		"",
		ddgSearch)
}
