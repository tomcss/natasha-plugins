package random

import (
	"fmt"
	"github.com/go-chat-bot/bot"
    "math/rand"
)

func genRandom(cmd *bot.Cmd) (msg string, err error) {

    if len( cmd.Args) == 0 {
        msg = fmt.Sprintf( "%d", rand.Int())
    } else if len( cmd.Args) == 1 {

        switch cmd.Args[0] {
            case "int":
                msg = fmt.Sprintf( "%d", rand.Int())
            case "uint":
                msg = fmt.Sprintf( "%d", rand.Uint32())
            case "float":
                msg = fmt.Sprintf( "%f", rand.Float32())
            default:
                msg = "Unknown type: "+cmd.Args[0]
        }
    } else {
        msg = "Invalid number of arguments"
    }

    return
}

func init() {
	bot.RegisterCommand(
		"rnd",
        "Generate a random number: rnd [type] [min max]",
		"",
		genRandom)
}
