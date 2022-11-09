package main

import (
	"rft/lambda2documentdb/internal/store"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event events.SQSEvent) error {
	for _, record := range event.Records {
		if err := store.Store(record.Body); err != nil {
			return err
		}
	}
	return nil
}
