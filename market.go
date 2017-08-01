package main

import (
	"github.com/yanzay/tbot"
)

func replyMarket(m *tbot.Message) {
	content := renderMarket()
	buttons := [][]string{
		{"Food", "Cloth", "Teck"},
		{"Back", "Home", "Help"},
	}
	m.ReplyKeyboard(content, buttons, tbot.WithMarkdown)
}

func renderMarket() string {
	content := "```\n"
	content += "Market"
	content += "```"
	return content
}
