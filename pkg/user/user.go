package user

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)
var (
	ErrorFailedToGetRecord = "Failed to get record"
	ErrorFailedToUnmarshalRecord = "Failed to marshal record"
)

type User struct(
	Email string `json:"email`
	Name string `json:"name"`
)

func GetUser(email, tableName string, dynamoDbClient dynamodbiface.DynamoDBAPI) (*User, error) {
	query := &dynamodb.GetItemInput{
		Key: map[string] * dynamodb.AttributeValue{
			"email": {
				S: aws.String(email)
			}
		},
		TableName: aws.String(tableName)
	}

	result, err := dynamoDbClient.GetItem(query)

	if err != nil {
		return nil, errors.New(ErrorFailedToGetRecord)
	}

	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)


	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

func GetAllUsers() {

}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() error {

}
