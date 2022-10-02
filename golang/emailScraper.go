package main

import (
	"context"
	"strings"

	"github.com/DusanKasan/parsemail"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Response struct {
	ContactEmail string `json:"contactEmail"`
	PhoneNumber  string `json:"phoneNumber"`
	NumPeople    string `json:"numPeople"`
}

type Email struct {
	doc *goquery.Document
}

func (e Email) contact() string {
	return e.findElementContent("body > div > div > div:nth-child(4) > div > table:nth-child(6) > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table > tbody > tr:nth-child(8) > td > table > tbody > tr:nth-child(7) > td > div > div > a")
}
func (e Email) phoneNumber() string {
	return e.findElementContent("body > div > div > div:nth-child(4) > div > table:nth-child(6) > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table > tbody > tr:nth-child(8) > td > table > tbody > tr:nth-child(7) > td > div > strong:nth-child(10) > a")
}
func (e Email) numPeople() string {
	return e.findElementContent("body > div > div > div:nth-child(4) > div > table:nth-child(6) > tbody > tr:nth-child(2) > td > table > tbody > tr > td > table > tbody > tr:nth-child(8) > td > table > tbody > tr:nth-child(10) > td > table > tbody > tr > td > table:nth-child(1) > tbody > tr:nth-child(1) > td:nth-child(1) > div")
}

func (e Email) findElementContent(selector string) string {
	return e.doc.Find(selector).Text()
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

	emailUtils := Email{doc: doc}
	contactEmail := emailUtils.contact()
	phoneNumber := emailUtils.phoneNumber()
	numPeople := emailUtils.numPeople()
	return &Response{ContactEmail: contactEmail, PhoneNumber: phoneNumber, NumPeople: numPeople}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
