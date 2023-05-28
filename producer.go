package kafka_client

import (
	"context"
)

type Producer struct {
	BaseKafka
}

func (t *Producer) Publish(ctx context.Context, data ...[]byte) error {
	for _, d := range data {
		err := t.publish(ctx, t.TopicConfig.Name, d)
		if err != nil {
			return err
		}
	}
	return nil
}
