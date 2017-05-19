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
	routerMux := tbot.NewRouterMux(tbot.NewSessionStorage())
	var s *tbot.Server
	var err error
	if *local {
		s, err = tbot.NewServer(token, tbot.WithMux(routerMux))
	} else {
		s, err = tbot.NewServer(token,
			tbot.WithMux(routerMux),
			tbot.WithWebhook("https://bot.yanzay.com/"+token, "0.0.0.0:8013"))
	}
	if err != nil {
		log.Fatal(err)
	}
	s.AddMiddleware(Login)
	s.HandleFunc(tbot.RouteRoot, HomeHandler)
	s.HandleFunc("/market", MarketHandler)
	s.SetAlias(tbot.RouteRoot, "Home")
	s.SetAlias(tbot.RouteBack, "Back")
	s.SetAlias("/market", "Market")
	s.HandleDefault(DefaultHandler)
	s.ListenAndServe()
}

func HomeHandler(m *tbot.Message) {
	ReplyHome(m)
}

func MarketHandler(m *tbot.Message) {
	ReplyMarket(m)
}

func DefaultHandler(m *tbot.Message) {
	m.Reply("hm?")
}
