package main

import (
	"os"

	"github.com/yanzay/log"
	"github.com/yanzay/tbot"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	log.Infof("Starting bot with token: %s", token)
	s, err := tbot.NewServer(token, tbot.WithWebhook("https://bot.yanzay.com/"+token, "0.0.0.0:8013"))
	if err != nil {
		log.Fatal(err)
	}
	s.HandleFunc("/start", StartHandler)
	s.ListenAndServe()
}

func StartHandler(m *tbot.Message) {
	m.Reply("Welcome!")
}
