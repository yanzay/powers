package main

import (
	"testing"

	"github.com/yanzay/powers/models"
	"github.com/yanzay/powers/storage"
	"github.com/yanzay/tbot"
	"github.com/yanzay/tbot/model"
)

func TestQuestion_isValidAnswer(t *testing.T) {
	type fields struct {
		Key               string
		Prompt            string
		Options           []string
		ValidationRule    string
		ValidationComment string
	}
	type args struct {
		answer string
	}
	questionName := questions["first_name"]
	questionAge := questions["18+"]
	tests := []struct {
		name   string
		fields *Question
		args   args
		want   string
		want1  bool
	}{
		{"valid name", questionName, args{"A"}, "", true},
		{"invalid name", questionName, args{"a"}, questionName.ValidationComment, false},
		{"empty name", questionName, args{""}, questionName.ValidationComment, false},
		{"name with spaces", questionName, args{"I am root"}, questionName.ValidationComment, false},

		{"18+ yes", questionAge, args{yesButton}, "", true},
		{"18+ no", questionAge, args{noButton}, "", true},
		{"incorrect age", questionAge, args{"yeah!"}, questionAge.ValidationComment, false},
		{"empty age", questionAge, args{""}, questionAge.ValidationComment, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.fields
			got, got1 := q.isValidAnswer(tt.args.answer)
			if got != tt.want {
				t.Errorf("Question.isValidAnswer() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Question.isValidAnswer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_survey(t *testing.T) {
	store = storage.New(".test.db")
	flow := []string{"/start", "Jon", "Snow", yesButton}
	want := []string{
		questions["first_name"].Prompt,
		questions["last_name"].Prompt,
		questions["18+"].Prompt,
		"Registered",
	}
	got := []string{}
	profile := &models.Profile{}
	for _, input := range flow {
		m := &tbot.Message{Message: &model.Message{}}
		replies := make(chan *model.Message, 100)
		m.SetReplyChannel(replies)
		m.Data = input
		survey(m, profile)
		reply := <-replies
		got = append(got, reply.Data)
	}
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("survey want %v, got %v", want[i], got[i])
		}
	}
}
