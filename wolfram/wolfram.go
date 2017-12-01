package wolframcmd

import (
    //"encoding/xml"
    "fmt"
    "net/http"
    "io/ioutil"
    "net/url"
    //"strings"
	"github.com/go-chat-bot/bot"
    //"github.com/Krognol/go-wolfram"
)

const APPID = "TV778Q-364AH8AV74"
const WOLFRAM_ENDPOINT = "http://api.wolframalpha.com/v1/result?appid="+APPID+"&i="

func wr(cmd *bot.Cmd) (msg string, err error) {

    //c := &wolfram.Client{AppID:"TV778Q-364AH8AV74x"}

    resp, err := http.Get( WOLFRAM_ENDPOINT + url.QueryEscape(cmd.RawArgs))

    if err != nil {
        msg = "Errorrrrr"
    } else {
        defer resp.Body.Close()
        b, _ := ioutil.ReadAll(resp.Body)

        msg = fmt.Sprintf("%s", b)
    }

    println(msg)

    return
}

func init() {
	bot.RegisterCommand(
		"calc",
		"Use Wolfram Alpha",
		"",
		wr)
}

