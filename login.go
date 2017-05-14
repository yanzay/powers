package main

import (
	"github.com/yanzay/log"
	"github.com/yanzay/tbot"
)

type Question struct {
	Key               string
	Prompt            string
	Options           []string
	ValidationRule    string
	ValidationComment string
	asking            bool
}

var questions = []*Question{
	{Key: "first_name", Prompt: "Enter your first name:"},
	{Key: "last_name", Prompt: "Enter your last name:"},
	{Key: "18+", Prompt: "Are you 18+ years old?", Options: []string{"Yes", "No"}},
}

func Login(f tbot.HandlerFunction) tbot.HandlerFunction {
	return func(m *tbot.Message) {
		log.Debugf("ChatID: %d", m.ChatID)
		if registered(m.ChatID) {
			f(m)
			return
		}
		questionare(m)
	}
}

func questionare(m *tbot.Message) {
	profile := storage.GetProfile(m.ChatID)
	checkAsking(profile, m)
	for _, question := range questions {
		if profile[question.Key] == "" {
			question.asking = true
			if question.Options != nil {
				m.ReplyKeyboard(question.Prompt, [][]string{question.Options}, tbot.OneTimeKeyboard)
			} else {
				m.Reply(question.Prompt)
			}
			return
		}
	}
	m.Reply("Registered")
}

func checkAsking(profile Profile, m *tbot.Message) {
	for _, question := range questions {
		if question.asking {
			profile[question.Key] = m.Text()
			question.asking = false
			return
		}
	}
}

func registered(chatID int64) bool {
	profile := storage.GetProfile(chatID)
	for _, question := range questions {
		if profile[question.Key] == "" {
			return false
		}
	}
	return true
}
