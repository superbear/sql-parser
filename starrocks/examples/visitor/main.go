package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(in prompt.Document) []prompt.Suggest {
	var ret []prompt.Suggest
	return ret
}

func main() {
	var p *prompt.Prompt

	executor := func(s string) {
		sql := strings.TrimSpace(s)
		if sql == "" {
			log.Println("sql required")
			return
		} else if sql == "quit" || sql == "exit" {
			exit()
			return
		}

		columns, err := parse(sql)
		if err == nil {
			fmt.Println("columns: ", columns)
		} else {
			log.Println("parse err: ", err)
		}
	}

	kb := []prompt.KeyBind{
		{
			Key: prompt.ControlD,
			Fn: func(buf *prompt.Buffer) {
				exit()
			},
		},
	}
	p = prompt.New(
		executor,
		completer,
		prompt.OptionAddKeyBind(kb...),
	)

	fmt.Println("Please input sql")
	p.Run()
}

func exit() {
	fmt.Println("bye")
	os.Exit(0)
}
