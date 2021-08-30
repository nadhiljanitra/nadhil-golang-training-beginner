package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func InitSQSClient() *sqs.SQS {
	region := os.Getenv("SQS_REGION")
	endpoint := os.Getenv("SQS_ENDPOINT")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	}))

	client := sqs.New(sess)
	return client
}
