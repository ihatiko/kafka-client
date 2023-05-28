package kafka_client

import (
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func (c *Config) NewWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(c.Host...),
		Balancer:               &kafka.LeastBytes{},
		RequiredAcks:           WriterRequiredAcksDefault,
		MaxAttempts:            WriterMaxAttemptsDefault,
		ReadTimeout:            WriterReadTimeoutDefault,
		WriteTimeout:           WriterWriteTimeoutDefault,
		Compression:            CompressionDefault,
		Async:                  AsyncDefault,
		AllowAutoTopicCreation: c.AllowAutoTopicCreation,
	}
	if c.UseSSL {
		w.Transport = &kafka.Transport{
			TLS: c.NewTlsConfig(),
			SASL: plain.Mechanism{
				Username: c.Username,
				Password: c.Password,
			},
		}
	}

	return w
}
