package test

import (
	"io/ioutil"
	"rft/emailscraper/internal/email_utils"
	"strings"
	"testing"

	"github.com/DusanKasan/parsemail"
	"github.com/PuerkitoBio/goquery"
)

/**
type Email interface {
	Contact() string
	PhoneNumber() string
	BookingNumber() string
	NumPeople() string
	Name() string
	Tour() string
	Date() string
}
**/

func TestContactField(t *testing.T) {
	email := email_utils.CreateEmailParser(createNewDocument(readHTMLBody()))
	want := "manolomanigua@hotmail.com"
	msg := email.Contact()
	if want != msg {
		t.Fatalf(`expected: %q, current: %q`, want, msg)
	}
}

func TestBookingNumberField(t *testing.T) {
	email := email_utils.CreateEmailParser(createNewDocument(readHTMLBody()))
	want := "B651-T1177-220919-1"
	msg := email.BookingNumber()
	if want != msg {
		t.Fatalf(`expected: %q, current: %q`, want, msg)
	}
}
func TestPhoneNumberField(t *testing.T) {
	email := email_utils.CreateEmailParser(createNewDocument(readHTMLBody()))
	want := "+579999999999"
	msg := email.PhoneNumber()
	if want != msg {
		t.Fatalf(`expected: %q, current: %q`, want, msg)
	}
}
func TestNumPeopleField(t *testing.T) {
	email := email_utils.CreateEmailParser(createNewDocument(readHTMLBody()))
	want := "22"
	msg := email.NumPeople()
	if want != msg {
		t.Fatalf(`expected: %q, current: %q`, want, msg)
	}
}
func TestNameField(t *testing.T) {
	email := email_utils.CreateEmailParser(createNewDocument(readHTMLBody()))
	want := "manolo manigua"
	msg := email.Name()
	if want != msg {
		t.Fatalf(`expected: %q, current: %q`, want, msg)
	}
}
func TestTourField(t *testing.T) {
	email := email_utils.CreateEmailParser(createNewDocument(readHTMLBody()))
	want := "Free Tour del Castillo, Catedral y Malá Strana"
	msg := email.Tour()
	if want != msg {
		t.Fatalf(`expected: %q, current: %q`, want, msg)
	}
}
func TestDateField(t *testing.T) {
	email := email_utils.CreateEmailParser(createNewDocument(readHTMLBody()))
	want := "14:30 15-10-2022"
	msg := email.Date()
	if want != msg {
		t.Fatalf(`expected: %q, current: %q`, want, msg)
	}
}

func readHTMLBody() *strings.Reader {
	content, err := ioutil.ReadFile("../assets/ue2solbogubgn87dsj38flqbv8dq6tttdoib92g1")
	if err != nil {
		panic(err)
	}
	contentStr := string(content)
	return strings.NewReader(contentStr)
}

func createNewDocument(reader *strings.Reader) *goquery.Document {
	emailContent, err := parsemail.Parse(reader) // returns Email struct and error
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(emailContent.HTMLBody))
	if err != nil {
		panic(err)
	}
	return doc
}
