package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "os"
    "fmt"
)

type Response events.APIGatewayProxyResponse
type Movie struct {
    Year   int
    Title  string
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	requestdata, err := json.Marshal(request)
	fmt.Println(string(requestdata))
	fmt.Println(request)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
    db := dynamodb.New(sess)

    tableName := os.Getenv("DYNAMO_TABLE_NAME_ONE")
    year := request.QueryStringParameters["year"]
    fmt.Println("tableName: ",tableName)
    fmt.Println("year: ",year)
    params := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditionExpression: aws.String("#yr = :yyyy"), // 'year' is reserved keyword
		ExpressionAttributeNames: map[string]*string{
			"#yr": aws.String("year"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":yyyy": {
				N: aws.String(year),
			},
		},
	}
	resp, err := db.Query(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return Response{StatusCode: 500}, err
	}

	// dump the response data
	fmt.Println(resp)

	// Unmarshal the slice of dynamodb attribute values
	// into a slice of custom structs
	var movies []Movie
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &movies)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

    responBody, _ := json.Marshal(movies)
	respon := Response{
		StatusCode: 200,
		IsBase64Encoded: false,
		Body: string(responBody),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "world-handler",
		},
	}

	return respon, nil
}

func main() {
	lambda.Start(Handler)
}
