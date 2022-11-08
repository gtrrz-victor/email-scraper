package dateparser

import (
	"fmt"
	"strings"
)

type month struct {
	spanish string
	val     string
}

type monthTranslator struct {
	months []month
}

func (mt monthTranslator) translate(spanishValue string) (string, error) {

	for _, month := range mt.months {
		if month.spanish == strings.ToLower(spanishValue) {
			return month.val, nil
		}
	}
	return "", fmt.Errorf("%v month can not be translated", spanishValue)

}
