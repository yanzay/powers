package main

import (
	"fmt"

	"github.com/yanzay/tbot"
)

func ReplyMarket(m *tbot.Message) {
	market := storage.GetMarket()
	content := renderMarket(market)
	buttons := [][]string{
		{"Food", "Cloth", "Teck"},
		{"Back", "Home", "Help"},
	}
	m.ReplyKeyboard(content, buttons, tbot.WithMarkdown)
}

func renderMarket(market *Market) string {
	content := "```\n"
	content += fmt.Sprint(market)
	content += "```"
	return content
}
