package cli

import (
	"strings"

	"github.com/ergochat/readline"
)

var rl *readline.Instance

//

func Init() error {
	var err error
	rl, err = readline.NewFromConfig(&readline.Config{Prompt: "> "})
	return err
}

func Ask(prompt string) (string, error) {
	if rl == nil {
		if err := Init(); err != nil {
			return "", err
		}
	}
	rl.SetPrompt(prompt)
	line, err := rl.Readline()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}
