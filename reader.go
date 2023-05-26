package kafka_client

import (
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"time"
)

func (c *Config) newReader(groupId string, topic ...string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:                c.Host,
		GroupID:                groupId,
		GroupTopics:            topic,
		MinBytes:               minBytesDefault,
		MaxBytes:               maxBytesDefault,
		QueueCapacity:          queueCapacityDefault,
		HeartbeatInterval:      heartbeatIntervalDefault,
		CommitInterval:         commitIntervalDefault,
		PartitionWatchInterval: partitionWatchIntervalDefault,
		MaxAttempts:            maxAttemptsDefault,
		MaxWait:                time.Second,
		Dialer:                 c.getDialer(),
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
			Timeout:       dialTimeoutDefault,
			DualStack:     true,
			TLS:           c.NewTlsConfig(),
			SASLMechanism: mechanism,
		}
	}
	return &kafka.Dialer{
		Timeout: dialTimeoutDefault,
	}
}
