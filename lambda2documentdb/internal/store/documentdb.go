package store

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"encoding/json"
	"fmt"
)

func Store(body string) error {
	fmt.Println("Message received. Body" + body)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	sqsMessage := &SQSWraperMessage{}
	if err := json.Unmarshal([]byte(body), sqsMessage); err != nil {
		return fmt.Errorf("invalid json format. Body: '%s'", body)
	}
	fmt.Printf("Unmarshal content received. Booking: %s", sqsMessage.ResponsePayload)

	tableName := "ExternalBookings"

	av, err := dynamodbattribute.MarshalMap(sqsMessage.ResponsePayload)
	if err != nil {
		return err
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}

	fmt.Println("Successfully added '" + body + "' to table " + tableName)
	return nil
}
