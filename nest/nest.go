package github.com.go-chat-bot.bot

import (
    "fmt"
    "github.com/go-chat-bot/bot"
    "regexp"
    //"github.com/whee/ddg"
    //"strings"
    //"net/url"
    //"github.com/PuerkitoBio/goquery"
)

const (
    pattern = `\[([^\[]+?)\]`
)

var (
    re = regexp.MustCompile( pattern)
)

    //func gif(command *bot.Cmd) (msg string, err error) {
func parseNest(cmd *bot.Cmd) (msg string, err error) {

    cmdLine := cmd.RawArgs

    if re.MatchString( cmdLine) {
        subCmd := re.FindStringSubmatch( cmdLine)[1]
        fmt.Println( subCmd)
        subResult, err := parse( subCmd, cmd.ChannelData, cmd.User)

        //m, e := message, err := subResult.command.CmdFuncV1(subResult)
        fmt.Println( subResult.Command, err)
        //cmdLine = re.ReplaceAllString( cmdLine, subResult.Message)
    }

    return cmdLine, nil
}

func init() {
    bot.RegisterCommand(
        "nest",
        "Execute nested commands (``)",
        "",
        parseNest)
}
