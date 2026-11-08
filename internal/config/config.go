package config
//
// import (
// 	"log"
// 	"strings"
//
// 	"github.com/ergochat/readline"
// )
//
// var rl *readline.Instance
//
// //
//
// func Init() {
// 	var err error
// 	rl, err = readline.NewFromConfig(&readline.Config{
// 		Prompt:      "> ",
// 		HistoryFile: "/tmp/gust_history",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
//
// func ask(prompt string) (string, error) {
// 	rl.SetPrompt(prompt)
// 	line, err := rl.Readline()
// 	if err != nil {
// 		return "", err
// 	}
// 	return strings.TrimSpace(line), nil
// }
