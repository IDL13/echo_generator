package parser

import (
	"flag"
	"fmt"
	"reflect"
)

type Parser struct {
	name  string
	value string
	usage string
}

func NewParser(data interface{}) ([]*bool, []*string) {
	var request []*Parser
	reflection := reflect.TypeOf(data)
	l := reflection.NumField()

	for i := 0; i < l; i++ {
		var p Parser

		p.name = reflection.Field(i).Tag.Get("name")
		p.usage = reflection.Field(i).Tag.Get("usage")
		p.value = reflection.Field(i).Tag.Get("value")
		request = append(request, &p)
	}
	var flagsBool []*bool
	var flagsString []*string

	for i := range request {
		if request[i].value == "false" || request[i].value == "true" {
			f := flag.Bool(request[i].name, boolDecorate(request[i].value), request[i].usage)
			flagsBool = append(flagsBool, f)
		} else {
			f := flag.String(request[i].name, StringDecorate(request[i].value), request[i].usage)
			flagsString = append(flagsString, f)
		}
	}

	return flagsBool, flagsString
}

func boolDecorate(b string) bool {
	if b == "true" {
		return true
	} else {
		return false
	}
}

func StringDecorate(s any) string {
	str := fmt.Sprintf("%s", s)
	return str
}
