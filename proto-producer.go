package kafka_client

import (
	"context"
	"google.golang.org/protobuf/proto"
)

type ProtoProducer[T proto.Message] struct {
	BaseProducer
}

func (t ProtoProducer[T]) Publish(ctx context.Context, data ...T) error {
	for _, d := range data {
		m, err := proto.Marshal(d)
		if err != nil {
			return err
		}
		err = t.publish(ctx, m)
		if err != nil {
			return err
		}
	}
	return nil
}
