package kafka_client

import (
	"context"
	jsoniter "github.com/json-iterator/go"
)

type ConsumerJson[T any] struct {
	BaseKafka
}

func (c *ConsumerJson[T]) Consume(handler Handler[T]) {
	ctx := context.Background()
	for _, rd := range c.Readers {
		go c.consume(ctx, rd, func(context context.Context, bytes []byte) error {
			data := new(T)
			err := jsoniter.Unmarshal(bytes, data)
			if err != nil {
				return err
			}
			return handler(&Request[T]{
				Data:    *data,
				Context: context,
			})
		})
	}
}
