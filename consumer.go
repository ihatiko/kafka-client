package kafka_client

import (
	"context"
)

const (
	attemptKey        = "attempt"
	replaceAttemptKey = ".attempt"
	waitForKey        = "waitFor"
	DLQKey            = "DLQ"
)

type Consumer struct {
	BaseKafka
}

func (c *Consumer) Consume(handler Handler[Data]) {
	ctx := context.Background()
	for _, rd := range c.Readers {
		go c.consume(ctx, rd, func(context context.Context, bytes []byte) error {
			return handler(&Request[Data]{
				Data:    bytes,
				Context: context,
			})
		})
	}
}
