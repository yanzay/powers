package main

import (
	"fmt"

	"github.com/yanzay/tbot"
)

func ReplyHome(m *tbot.Message) {
	home := storage.GetHome(m.ChatID)
	content := renderHome(home)
	m.ReplyKeyboard(content, [][]string{{"Home", "Market"}}, tbot.WithMarkdown)
}

func renderHome(home *Home) string {
	content := "```\n"
	content += fmt.Sprintf("Money: %d\n", home.Money)
	content += fmt.Sprintf("Size: %d\n", home.Size)
	content += "```"
	return content
}
