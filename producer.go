package kafka_client

import (
	"context"
)

type Producer struct {
	BaseProducer
}

func (t *Producer) Publish(ctx context.Context, data ...[]byte) error {
	for _, d := range data {
		err := t.publish(ctx, d)
		if err != nil {
			return err
		}
	}
	return nil
}
