package aws

import (
	"context"
)

type (
	MessageClient interface {
		// CreateQueue Creates a new long polling queue and returns its URL.
		CreateQueue(ctx context.Context, queueName string, isDLX bool) (string, error)
		// QueueARN Get a queue ARN.
		QueueARN(ctx context.Context, queueURL string) (string, error)
		// BindDLX Binds a DLX queue to a normal queue.
		BindDLX(ctx context.Context, queueURL, dlxARN string) error
		// SendMessage Send a message to queue and returns its message ID.
		SendMessage(ctx context.Context, req *SendRequest) (string, error)
		// ReceiveMessage Long polls given amount of messages from a queue.
		ReceiveMessage(ctx context.Context, queueURL string) (*Message, error)
		// DeleteMessage Deletes a message from a queue.
		DeleteMessage(ctx context.Context, queueURL, rcvHandle string) error
	}
)