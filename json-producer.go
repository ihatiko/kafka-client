package kafka_client

import (
	"context"
	jsoniter "github.com/json-iterator/go"
)

type JsonProducer[T any] struct {
	BaseProducer
}

func (t JsonProducer[T]) Publish(ctx context.Context, data ...T) error {
	for _, d := range data {
		m, err := jsoniter.Marshal(d)
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
