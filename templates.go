package main

import (
	"strings"
	"text/template"
	"unicode/utf8"
)

const padWidth = 23

var funcMap = template.FuncMap{"pad": pad}
var homeTmpl = template.Must(template.New("home").Funcs(funcMap).Parse(homeTemplate))

func pad(first string, last string) string {
	firstLen := utf8.RuneCountInString(first)
	lastLen := utf8.RuneCountInString(last)
	if firstLen+lastLen > padWidth-1 {
		r := []rune(first)
		r = r[:padWidth-lastLen-1]
		first = string(r)
		firstLen = utf8.RuneCountInString(first)
	}
	repeatCount := padWidth - firstLen - lastLen
	return first + strings.Repeat(".", repeatCount) + last
}
