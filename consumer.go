package kafka_client

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	BaseKafka
}

func (c *Consumer) Consume(handler Handler[Data]) {
	go func() {
		//TODO shutdown
		ctx := context.Background()
		for {
			m, err := c.Reader.FetchMessage(ctx)
			if err != nil {
				break
			}
			extractedContext := ExtractJaegerContext(m)
			request := Request[Data]{
				Data:    m.Value,
				Context: extractedContext,
			}
			err = handler(&request)
			if err != nil && c.ConsumerConfig.DLQ != nil {
				if c.Writer == nil {
					c.Writer = c.KafkaConfig.NewWriter()
				}
				if extractedContext == nil {
					extractedContext = context.TODO()
				}
				err = c.publish(
					extractedContext,
					fmt.Sprintf("%s.attempts.1", m.Topic),
					m.Value,
					kafka.Header{
						Key:   "attempt",
						Value: []byte("1"),
					})
				if err != nil {
					log.Fatal("failed to commit dlq message:", err)
				}
			}
			if err := c.Reader.CommitMessages(ctx, m); err != nil {
				log.Fatal("failed to commit messages:", err)
			}
		}
	}()
}
