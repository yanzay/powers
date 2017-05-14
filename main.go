package main

import (
	"flag"
	"os"

	"github.com/yanzay/log"
	"github.com/yanzay/tbot"
)

var (
	local = flag.Bool("local", false, "Launch bot without webhook")
)

var storage = NewStorage()

func main() {
	flag.Parse()
	token := os.Getenv("TELEGRAM_TOKEN")
	log.Infof("Starting bot with token: %s", token)
	treeMux := &TreeMux{}
	var s *tbot.Server
	var err error
	if *local {
		s, err = tbot.NewServer(token, tbot.WithMux(treeMux))
	} else {
		s, err = tbot.NewServer(token,
			tbot.WithMux(treeMux),
			tbot.WithWebhook("https://bot.yanzay.com/"+token, "0.0.0.0:8013"))
	}
	if err != nil {
		log.Fatal(err)
	}
	s.AddMiddleware(Login)
	s.HandleFunc("/start", StartHandler)
	s.HandleDefault(DefaultHandler)
	s.ListenAndServe()
}

func StartHandler(m *tbot.Message) {
	m.Reply("Welcome to the Powers!")
}

func DefaultHandler(m *tbot.Message) {
	m.Reply("hm?")
}
