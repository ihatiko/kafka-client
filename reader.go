package kafka_client

import (
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func (c *Config) newReader(group *ConsumerGroup) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:                c.Host,
		GroupID:                group.GroupID,
		GroupTopics:            group.Topics,
		MinBytes:               MinBytes,
		MaxBytes:               MaxBytes,
		QueueCapacity:          QueueCapacity,
		HeartbeatInterval:      HeartbeatInterval,
		CommitInterval:         CommitInterval,
		PartitionWatchInterval: PartitionWatchInterval,
		MaxAttempts:            MaxAttempts,
		MaxWait:                MaxWait,
		Dialer:                 c.getDialer(),
		//TODO добавить оставшиеся параметры для переопределения
	})

	return r
}

func (c *Config) getDialer() *kafka.Dialer {
	if c.UseSSL {
		mechanism := plain.Mechanism{
			Username: c.Username,
			Password: c.Password,
		}
		return &kafka.Dialer{
			Timeout:       DialTimeout,
			DualStack:     true,
			TLS:           c.NewTlsConfig(),
			SASLMechanism: mechanism,
		}
	}
	return &kafka.Dialer{
		Timeout: DialTimeout,
	}
}
