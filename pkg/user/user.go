package user

import (
	"encoding/json"
	"errors"
	"github.com/benmotyka/boring-serverless-crud/pkg/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)
var (
	ErrorFailedToGetRecord = "Failed to get record"
	ErrorFailedToUnmarshalRecord = "Failed to unmarshal record"
	ErrorInvalidUserData = "Invalid user data"
	ErrorInvalidItem = "Invalid item"
	ErrorItemExists = "Item exists"
	ErrorCouldNotCreateItem = "Could not create item"
)

type User struct{
	Email string `json:"email`
	Name string `json:"name"`
}

func GetUser(email, tableName string, dynamoDbClient dynamodbiface.DynamoDBAPI) (*User, error) {
	query := &dynamodb.GetItemInput{
		Key: map[string] * dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynamoDbClient.GetItem(query)

	if err != nil {
		return nil, errors.New(ErrorFailedToGetRecord)
	}

	result := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, result)


	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return result, nil
}

func GetAllUsers(tableName string, dynamoDbClient dynamodbiface.DynamoDBAPI)(*[]User, error) {
query := &dynamodb.ScanInput{
	TableName: aws.String(tableName)
}

//scan = findall
result, err := dynamoDbClient.Scan(query)
if err != nil {
	return nil, errors.New(ErrorFailedToGetRecord)
}

result := new([]User)
err = dynamodbattribute.UnmarshalMap(result.Items, result)
return result, nil

}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynamoDbClient dynamodbiface.DynamoDBAPI) (*User, error){

	var user User 

	if err := json.Unmarshal([]byte(req.body), &user); err!=nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	if !validators.IsEmailValid(user.Email) {
		return nil, errors.New(ErrorInvalidUserData)
	}

	existingUser, err := GetUser(user.Email, tableName, dynamoDbClient)
	if existingUser != nil && len(existingUser.Email) != 0 {
		return nil, errors.New(ErrorItemExists)
	}

	item, err := dynamodbattribute.MarshalMap(user)

	if err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	query := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(tableName)
	}

	_, err = dynamoDbClient.PutItem(input)

	if err != nil {
		return nil, errors.New(ErrorCouldNotCreateItem)
	}

	return &user, nil
}

func UpdateUser() {

}

func DeleteUser() error {

}
