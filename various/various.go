package various

import (
	"github.com/go-chat-bot/bot"
)

func stabble(cmd *bot.Cmd) (msg string, err error) {

    if len( cmd.Args) == 0 {
        msg = "\u0001ACTION stabbles no-one. :(\u0001"
    } else {
        msg = "\u0001ACTION stabbles "+cmd.RawArgs+".\u0001"
    }

    return
}

func hugglefuck(cmd *bot.Cmd) (msg string, err error) {

    if len( cmd.Args) == 0 {
        msg = "\u0001ACTION hugglefucks no-one. :(\u0001"
    } else {
        msg = "\u0001ACTION hugglefucks "+cmd.RawArgs+" <3\u0001"
    }

    return
}

func newcmd(cmd *bot.Cmd) (msg string, err error) {

    if len( cmd.Args) == 0 {
        msg = "no command specified"
    } else if len( cmd.Args) == 1 {
        msg = "no action specified"
    } else {
        args := cmd.Args

	    bot.RegisterCommand(
		    cmd.Args[0],
            "made",
		    "",
            func(cmd2 *bot.Cmd) (msg string, err error) {
                msg = args[1]
                return msg, err
            })
    }

    return
}

func init() {
	bot.RegisterCommand(
		"stabble",
        "stabbles someone",
		"",
		stabble)

	bot.RegisterCommand(
		"hf",
        "hugglefucks someone",
		"",
		hugglefuck)

    bot.RegisterCommand(
        "new",
        "",
        "",
        newcmd)
}
