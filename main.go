package main

import (
	"log"
	"os"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var c = metricsClient()

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if err := c.Incr("lambda_test", []string{}, 1); err != nil {
		log.Println("error while sending metric: ", err)
	}

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func metricsClient() *statsd.Client {
	dogstatsd_client, err := statsd.New(os.Getenv("DATADOG_AGENT"))
	if err != nil {
		log.Fatal("error while creating metrics client", err)
	}

	return dogstatsd_client
}

func main() {
	lambda.Start(Handler)
}
