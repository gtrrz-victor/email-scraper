package main

import (
	"rft/bookings/internal/get_bookings"
	"rft/bookings/internal/models"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest() ([]models.Response, error) {
	return get_bookings.Query()
}
