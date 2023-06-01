package kafka_client

import (
	"context"
	"github.com/segmentio/kafka-go"
)

const (
	attemptKey   = "attempt"
	waitForKey   = "wait-for"
	dLQKey       = "DLQ"
	mainTopicKey = "main-topic"
)

type Consumer struct {
	BaseKafka
}

func (c *Consumer) Consume(handler Handler[Data]) {
	ctx := context.Background()
	for _, rd := range c.Readers {
		go c.consume(ctx, rd, func(ctx context.Context, headers []kafka.Header, bytes []byte, s string, i int, i2 int64) error {
			return handler(&Request[Data]{
				Data:    bytes,
				Context: ctx,
			})
		})
	}
}
