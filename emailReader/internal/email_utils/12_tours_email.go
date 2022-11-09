package email_utils

import (
	"fmt"
	"regexp"
	"rft/emailscraper/internal/dateparser"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type oneTwoToursEmail struct {
	doc *goquery.Document
}

func (e oneTwoToursEmail) Contact() string {
	return e.findFirstElementContent("body table tbody div div a")
}
func (e oneTwoToursEmail) PhoneNumber() string {
	return e.findFirstElementContent("body table div strong a")
}
func (e oneTwoToursEmail) NumPeople() string {
	rawValue := e.findFirstElementContent("body table table table table table tr:nth-child(1) > td:nth-child(1) > div")
	re := regexp.MustCompile("[0-9]+")
	return re.FindString(rawValue)
}
func (e oneTwoToursEmail) Name() string {
	rawValue := e.findFirstElementContent("body table tbody tr > td > div > strong:nth-child(6)")
	return strings.Join(strings.Fields(strings.TrimSpace(rawValue)), " ")
}
func (e oneTwoToursEmail) Tour() string {
	query := e.doc.Find("body table tbody tr:nth-child(2) > td > div > br:last-child")
	if query.Length() == 1 {
		rawValue := query.Get(0).NextSibling.Data
		return strings.Join(strings.Fields(strings.TrimSpace(strings.ReplaceAll(rawValue, "\n", ""))), " ")
	}
	return "ERROR: Tour name could not be found"
}
func (e oneTwoToursEmail) BookingNumber() string {
	return e.findFirstElementContent("body table td div b")
}

func (e oneTwoToursEmail) Date() string {
	dateElement := e.findFirstElementContent("body table td div span:nth-child(4)")
	day, month, year, time := dateparser.ParseStringDate(strings.TrimSpace(dateElement))
	return fmt.Sprintf("%v %v-%v-%v", time, day, month, year)
}

func (e oneTwoToursEmail) findFirstElementContent(selector string) string {
	return e.doc.Find(selector).First().Text()
}
