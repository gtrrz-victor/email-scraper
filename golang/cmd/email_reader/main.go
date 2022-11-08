package main

import (
	"context"
	"rft/emailscraper/internal/email_utils"
	"strings"

	"github.com/DusanKasan/parsemail"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Response struct {
	ContactEmail  string `json:"contactEmail"`
	PhoneNumber   string `json:"phoneNumber"`
	NumPeople     string `json:"numPeople"`
	Date          string `json:"date"`
	Name          string `json:"name"`
	BookingNumber string `json:"bookingNumber"`
	Tour          string `json:"tour"`
	RawEmail      string `json:"rawemail"`
}

func HandleRequest(ctx context.Context, event events.S3Event) (*Response, error) {
	bucket := event.Records[0].S3.Bucket.Name
	key := event.Records[0].S3.Object.Key

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	s3Object, err := svc.GetObject(&s3.GetObjectInput{Bucket: &bucket, Key: &key})
	if err != nil {
		return nil, err
	}
	email, err := parsemail.Parse(s3Object.Body) // returns Email struct and error
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(email.HTMLBody))
	if err != nil {
		return nil, err
	}

	emailUtils := email_utils.CreateEmailParser(doc)
	contactEmail := emailUtils.Contact()
	phoneNumber := emailUtils.PhoneNumber()
	numPeople := emailUtils.NumPeople()
	date := emailUtils.Date()
	name := emailUtils.Name()
	bookingNumber := emailUtils.BookingNumber()
	tour := emailUtils.Tour()
	return &Response{ContactEmail: contactEmail, PhoneNumber: phoneNumber, NumPeople: numPeople, Date: date, Name: name, BookingNumber: bookingNumber, Tour: tour, RawEmail: key}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
