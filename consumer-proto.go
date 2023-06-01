package kafka_client

import (
	"context"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type ConsumerProto[T proto.Message] struct {
	BaseKafka
}

func (c *ConsumerProto[T]) Consume(handler Handler[T]) {
	ctx := context.Background()
	for _, rd := range c.Readers {
		go c.consume(ctx, rd, func(ctx context.Context, headers []kafka.Header, bytes []byte, s string, i int, i2 int64) error {
			data := new(T)
			err := proto.Unmarshal(bytes, *data)
			if err != nil {
				return err
			}
			return handler(&Request[T]{
				Data:    *data,
				Context: ctx,
			})
		})
	}
}
