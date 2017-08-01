package main

import (
	"regexp"

	"github.com/yanzay/log"
	"github.com/yanzay/powers/models"
	"github.com/yanzay/tbot"
)

type Question struct {
	Key               string
	Prompt            string
	Options           []string
	ValidationRule    string
	ValidationComment string
}

func (q *Question) isValidAnswer(answer string) (string, bool) {
	if q.ValidationRule != "" {
		match, err := regexp.MatchString(q.ValidationRule, answer)
		if err != nil {
			log.Errorf("error matching validation rule: %q", err)
		}
		if !match {
			return q.ValidationComment, false
		}
	}
	return "", true
}

var questions = map[string]*Question{
	"first_name": {
		Key:               "first_name",
		Prompt:            "Enter your first name:",
		ValidationRule:    "^[A-Z][a-z]*$",
		ValidationComment: "First name should start with capital letter and contain only letters",
	},
	"last_name": {
		Key:               "last_name",
		Prompt:            "Enter your last name:",
		ValidationRule:    "^[A-Z][a-z]*$",
		ValidationComment: "Last name should start with capital letter and contain only letters",
	},
	"18+": {
		Key:            "18+",
		Prompt:         "Are you 18+ years old?",
		ValidationRule: "Yes|No",
		Options:        []string{"Yes", "No"},
	},
}

func login(f tbot.HandlerFunction) tbot.HandlerFunction {
	return func(m *tbot.Message) {
		if registered(m.ChatID) {
			f(m)
			return
		}
		survey(m)
	}
}

func survey(m *tbot.Message) {
	profile, err := store.GetProfile(m.ChatID)
	if err != nil {
		log.Errorf("can't get player's profile: %q", err)
		return
	}
	survey, err := store.GetSurvey("login", m.ChatID)
	if err != nil {
		log.Errorf("can't get survey: %q", err)
		return
	}
	if survey.Asking != "" {
		comment := setAnswer(profile, questions[survey.Asking], m.Text())
		if comment != "" {
			m.Reply(comment)
			return
		}
		survey.Asking = ""
		store.SetSurvey("login", m.ChatID, survey)
		store.SetProfile(m.ChatID, profile)
	}
	if !registered(m.ChatID) {
		survey.Asking = askNext(profile, m)
		err = store.SetSurvey("login", m.ChatID, survey)
		if err != nil {
			log.Errorf("can't save survey %s: %q", "login", err)
		}
		return
	}

	if registered(m.ChatID) {
		m.Reply("Registered")
		replyHome(m)
	}
}

func askNext(profile *models.Profile, m *tbot.Message) string {
	question := nextQuestion(profile)
	if question == nil {
		return ""
	}
	if question.Options != nil {
		m.ReplyKeyboard(question.Prompt, [][]string{question.Options}, tbot.OneTimeKeyboard)
	} else {
		m.Reply(question.Prompt)
	}
	return question.Key
}

func nextQuestion(profile *models.Profile) *Question {
	switch {
	case profile.FirstName == "":
		return questions["first_name"]
	case profile.LastName == "":
		return questions["last_name"]
	case profile.Has18 == false:
		return questions["18+"]
	}
	return nil
}

func setAnswer(prof *models.Profile, question *Question, answer string) string {
	if comment, ok := question.isValidAnswer(answer); !ok {
		return comment
	}
	switch question.Key {
	case "first_name":
		prof.FirstName = answer
	case "last_name":
		prof.LastName = answer
	case "18+":
		prof.Has18 = (answer == "Yes")
	}
	return ""
}

func registered(chatID int64) bool {
	profile, err := store.GetProfile(chatID)
	if err != nil {
		log.Errorf("can't get player's profile: %q", err)
		return false
	}
	return profile.IsFull()
}
