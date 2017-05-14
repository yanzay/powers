package main

import "github.com/yanzay/tbot"

type TreeMux struct{}

func (tm *TreeMux) Mux(string) (*tbot.Handler, tbot.MessageVars) {
	return &tbot.Handler{}, tbot.MessageVars{}
}

func (tm *TreeMux) HandleFunc(string, tbot.HandlerFunction, ...string) {
}

func (tm *TreeMux) HandleFile(tbot.HandlerFunction, ...string)    {}
func (tm *TreeMux) HandleDefault(tbot.HandlerFunction, ...string) {}

func (tm *TreeMux) Handlers() tbot.Handlers {
	return tbot.Handlers{}
}

func (tm *TreeMux) DefaultHandler() *tbot.Handler {
	return &tbot.Handler{}
}

func (tm *TreeMux) FileHandler() *tbot.Handler {
	return &tbot.Handler{}
}
