package kafka

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

// MessageHandler represents a message handler.
type MessageHandler func(kgo.FetchTopicPartition)

// Client represents a Kafka client.
type Client struct {
	Kgo *kgo.Client
}

// NewClient creates a new Kafka client.
func NewClient(opt ...kgo.Opt) (*Client, error) {
	c, err := kgo.NewClient(opt...)
	if err != nil {
		return nil, err
	}
	return &Client{
		Kgo: c,
	}, nil
}

// ProduceMessage produces a message to a Kafka topic.
func (c *Client) ProduceMessage(ctx context.Context, rs ...*kgo.Record) error {
	if err := c.Kgo.ProduceSync(ctx, rs...).FirstErr(); err != nil {
		return fmt.Errorf("record had a produce error while synchronously producing: %w", err)
	}
	return nil
}

// ConsumeMessage from a Kafka topic.
func (c *Client) PollFetches(ctx context.Context) (kgo.Fetches, error) {
	fetch := c.Kgo.PollFetches(ctx)

	if err := fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetch had an error: %w", err)
	}

	return fetch, nil
}

// ConsumeMessage from a Kafka topic.
func (c *Client) ConsumeMessage(ctx context.Context, fetches kgo.Fetches, fn MessageHandler) error {
	fetches.EachPartition(fn)
	if err := c.Kgo.CommitRecords(ctx, fetches.Records()...); err != nil {
		return fmt.Errorf("committing records failed: %w", err)
	}
	return nil
}

// Close the Kafka client.
func (c *Client) Close() {
	c.Kgo.Close()
}
