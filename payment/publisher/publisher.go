package publisher

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"golang.org/x/net/context"
)

type Publisher interface {
	Publish(ctx context.Context, payload []byte) error
}

type SqsPublisher struct {
	client            *sqs.SQS
	queueURL          string
	visibilityTimeout time.Duration
}

func NewSQSPublisher(client *sqs.SQS, queueURL string) Publisher {
	if client == nil {
		panic("sqs client is required") // must never happen
	}
	if len(queueURL) <= 0 {
		panic("queueURL can't be empty")
	}

	return SqsPublisher{
		queueURL:          queueURL,
		client:            client,
		visibilityTimeout: 30 * time.Second,
	}
}

func (s SqsPublisher) Publish(ctx context.Context, payload []byte) error {
	_, err := s.client.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(payload)),
		QueueUrl:    aws.String(s.queueURL),
	})
	if err != nil {
		return err
	}

	return nil
}
