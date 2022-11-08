package dateparser

import (
	"strings"
)

var translator monthTranslator

func init() {
	translator = monthTranslator{
		[]month{
			{spanish: "enero", val: "1"},
			{spanish: "febrero", val: "2"},
			{spanish: "marzo", val: "3"},
			{spanish: "abril", val: "4"},
			{spanish: "mayo", val: "5"},
			{spanish: "junio", val: "6"},
			{spanish: "julio", val: "7"},
			{spanish: "agosto", val: "8"},
			{spanish: "septiembre", val: "9"},
			{spanish: "octubre", val: "10"},
			{spanish: "noviembre", val: "11"},
			{spanish: "diciembre", val: "12"},
		},
	}
}

func ParseStringDate(date string) (string, string, string, string) {
	values := strings.Split(date, " ")[1:]
	return values[0], translateMonth(values[1]), values[2], values[3]
}

func translateMonth(monthSpanish string) string {
	if val, err := translator.translate(monthSpanish); err != nil {
		return monthSpanish
	} else {
		return val
	}
}
