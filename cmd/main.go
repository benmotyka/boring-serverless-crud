package main

import(
"github.com/benmotyka/serverless-crud/pkg/handlers"
"os"
"github.com/aws/aws-lambda-go/events"
"github.com/aws/aws-lambda-go/lambda"
"github.com/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/dynamodb"
"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var(
	dynamoDbClient dynamodbiface.DynamoDBAPI
)

func main(){
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.newSession(&aws.Config{
		Region: aws.String(region)
		}
	)
	if err != nil {
		return
	}
	dynamoDbClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = 'BoringServerlessCrud'

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error){
	switch req.HTTPMethod{
	case "GET":
			return handlers.GetUser(req, tableName, dynamoDbClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynamoDbClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynamoDbClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynamoDbClient)
	}
	default:
		return handlers.UnhandledMethod()
}