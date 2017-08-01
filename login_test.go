package main

import (
	"reflect"
	"testing"

	"github.com/yanzay/powers/models"
	"github.com/yanzay/tbot"
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
	questionName := fields{"name", "name?", nil, "^[A-Z][a-z]*$", "name should start with capital letter"}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		{"valid name", questionName, args{"A"}, "", true},
		{"invalid name", questionName, args{"a"}, questionName.ValidationComment, false},
		{"empty name", questionName, args{""}, questionName.ValidationComment, false},
		{"name with spaces", questionName, args{"I am root"}, questionName.ValidationComment, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Question{
				Key:               tt.fields.Key,
				Prompt:            tt.fields.Prompt,
				Options:           tt.fields.Options,
				ValidationRule:    tt.fields.ValidationRule,
				ValidationComment: tt.fields.ValidationComment,
			}
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

func Test_login(t *testing.T) {
	type args struct {
		f tbot.HandlerFunction
	}
	tests := []struct {
		name string
		args args
		want tbot.HandlerFunction
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := login(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_survey(t *testing.T) {
	type args struct {
		m       *tbot.Message
		profile *models.Profile
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			survey(tt.args.m, tt.args.profile)
		})
	}
}

func Test_askNext(t *testing.T) {
	type args struct {
		profile *models.Profile
		m       *tbot.Message
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := askNext(tt.args.profile, tt.args.m); got != tt.want {
				t.Errorf("askNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextQuestion(t *testing.T) {
	type args struct {
		profile *models.Profile
	}
	tests := []struct {
		name string
		args args
		want *Question
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextQuestion(tt.args.profile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setAnswer(t *testing.T) {
	type args struct {
		prof     *models.Profile
		question *Question
		answer   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setAnswer(tt.args.prof, tt.args.question, tt.args.answer); got != tt.want {
				t.Errorf("setAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}
