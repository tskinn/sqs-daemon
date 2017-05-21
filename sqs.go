package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Something good
type SQS struct {
	cfg        Config
	sqsService *sqs.SQS
}

// Get the next batch of messages from the queue
func (instance SQS) NextMessages(queue string) ([]*sqs.Message, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(instance.cfg.SQSURL),
		MaxNumberOfMessages: aws.Int64(cfg.Connections),
		WaitTimeSeconds:     aws.Int64(instance.cfg.WaitTime),
	}

	resp, err := instance.sqsService.ReceiveMessage(params)
	if err != nil {
		return nil, err
	}
	return resp.Messages, nil
}

// Delete the message from the queue
func (instance SQS) Complete(queue, messageID string) (err error) {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(instance.cfg.SQSURL),
		ReceiptHandle: aws.String(messageID),
	}

	_, err = instance.sqsService.DeleteMessage(params)
	if err != nil {
		return err
	}

	return nil
}
