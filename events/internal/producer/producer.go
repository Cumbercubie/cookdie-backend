package producer

import (
	"context"
	events "cookdie/events/internal/models"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Queue interface {
	Publish(ctx context.Context, event events.Event) error
	Poll() ([]events.Event, error)
}

type SQSProducer struct {
	client   *sqs.Client
	queueURL string
}

type Producer struct {
	queue Queue
}

func NewSQSProducer(cfg aws.Config, queueURL string) *SQSProducer {
	return &SQSProducer{
		client:   sqs.NewFromConfig(cfg),
		queueURL: queueURL,
	}
}
func (q *SQSProducer) Publish(ctx context.Context, event events.Event) error {
	data, err := json.Marshal(event)

	if err != nil {
		return err
	}

	_, err = q.client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody:  aws.String(string(data)),
		QueueUrl:     &q.queueURL,
		DelaySeconds: 0,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"event_type": {
				DataType:    aws.String("string"),
				StringValue: aws.String(string(event.Type)),
			},
			"business_id": {
				DataType:    aws.String("string"),
				StringValue: aws.String(event.BusinessId),
			},
		},
	})

	return err
}

func (q *SQSProducer) Poll() ([]events.Event, error) {
	output, err := q.client.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
		QueueUrl:            &q.queueURL,
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     20,
	})
	var polledEvents []events.Event

	if err != nil {
		return nil, fmt.Errorf("receive message failed: %w", err)
	}

	for i, msg := range output.Messages {
		var event events.Event
		if err := json.Unmarshal([]byte(*msg.Body), &event); err != nil {
			fmt.Printf("failed to unmarshal message %d: %v\n", i, err)
			continue
		}
		polledEvents = append(polledEvents, event)

	}

	return polledEvents, nil

}
