package email_utils

import "github.com/PuerkitoBio/goquery"

type Email interface {
	Contact() string
	PhoneNumber() string
	NumPeople() string
	Name() string
	Tour() string
	BookingNumber() string
	Date() string
}

func CreateEmailParser(doc *goquery.Document) Email {
	return oneTwoToursEmail{doc: doc}
}
