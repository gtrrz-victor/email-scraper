package get_bookings

import (
	"context"
	"rft/bookings/internal/models"

	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Query() ([]models.Response, error) {
	bookings := []models.Response{}
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "eu-west-1"
		return nil
	})

	if err != nil {
		return bookings, err
	}
	svc := dynamodb.NewFromConfig(cfg)

	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("ExternalBookings"),
	})

	if err != nil {
		return bookings, err
	}

	err = attributevalue.UnmarshalListOfMaps(out.Items, &bookings)
	if err != nil {
		return bookings, fmt.Errorf("failed to unmarshal Dynamodb Scan Items, %v", err)
	}
	return bookings, nil
}
