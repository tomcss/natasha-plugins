package rpncmd

import (
	"fmt"
	"os"
    "github.com/sauerbraten/rpncalc"
    "github.com/go-chat-bot/bot"
)

    //func gif(command *bot.Cmd) (msg string, err error) {
func rpnEval(cmd *bot.Cmd) (msg string, err error) {
	// os.Args[0] is the program path

    s := rpn.Eval(cmd.Args)
	for _, token := range cmd.Args[1:] {
		rpn.Eval(token)
	}

	// final result should be the only element now on the stack
	result, err := s.Pop()
	if err != nil {
		panic(err)
	}

    msg = result
	fmt.Println(result)

	// make sure stack is now empty
	_, err = s.Pop()
	if err == nil {
		fmt.Fprintln(os.Stderr, "stack not empty!")
	}
}

func init() {
	bot.RegisterCommand(
		"rpn",
		"Reversed Polish Number Calculator",
		"",
		rpnEval)
}
