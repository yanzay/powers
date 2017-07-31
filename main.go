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
	addHandlers(s)
	addAliases(s)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func addHandlers(s *tbot.Server) {
	s.HandleFunc(tbot.RouteRoot, homeHandler)
	s.HandleFunc("/market", marketHandler)
	s.HandleFunc("/work", workHandler)
	s.HandleDefault(defaultHandler)
}

func addAliases(s *tbot.Server) {
	s.SetAlias(tbot.RouteRoot, "Home")
	s.SetAlias(tbot.RouteBack, "Back")
	s.SetAlias("market", "Market")
}

func homeHandler(m *tbot.Message) {
	ReplyHome(m)
}

func marketHandler(m *tbot.Message) {
	ReplyMarket(m)
}

func defaultHandler(m *tbot.Message) {
	m.Reply("hm?")
}

func workHandler(m *tbot.Message) {
}
