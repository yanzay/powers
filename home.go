package main

import (
	"bytes"
	"fmt"

	"github.com/yanzay/log"
	"github.com/yanzay/powers/models"
	"github.com/yanzay/tbot"
)

const homeTemplate = `Home

{{pad "Money" (print .Money)}}
{{pad "Size" (print .Size)}}`

func replyHome(m *tbot.Message) {
	home, err := store.GetHome(m.ChatID)
	content, err := renderHome(home)
	if err != nil {
		log.Errorf("can't render home: %q", err)
		return
	}
	m.ReplyKeyboard(content, [][]string{{"Home", "Market"}}, tbot.WithMarkdown)
}

func renderHome(home *models.Home) (string, error) {
	w := &bytes.Buffer{}
	fmt.Fprint(w, "```\n")
	err := homeTmpl.Execute(w, home)
	if err != nil {
		return "", err
	}
	fmt.Fprint(w, "```")
	return w.String(), err
}
