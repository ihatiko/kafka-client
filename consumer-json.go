package kafka_client

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
)

type ConsumerJson[T any] struct {
	BaseKafka
}

func (c *ConsumerJson[T]) Consume(handler Handler[T]) {
	ctx := context.Background()
	for _, rd := range c.Readers {
		go c.consume(ctx, rd, func(ctx context.Context, headers []kafka.Header, bytes []byte, s string, i int, i2 int64) error {
			data := new(T)
			err := jsoniter.Unmarshal(bytes, data)
			if err != nil {
				//TODO error log
				return err
			}
			return handler(&Request[T]{
				Data:    *data,
				Context: ctx,
			})
		})
	}
}
