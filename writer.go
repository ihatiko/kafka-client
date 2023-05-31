package kafka_client

import (
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func (c *Config) NewWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(c.Host...),
		Balancer:               &kafka.LeastBytes{},
		RequiredAcks:           WriterRequiredAcks,
		ReadTimeout:            ReadTimeOut,
		WriteTimeout:           WriteTimeOut,
		Compression:            Compression,
		Async:                  AsyncDefault,
		MaxAttempts:            MaxAttempts,
		AllowAutoTopicCreation: AllowAutoTopicCreation,
	}

	c.setEnvironments(w)
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

func (c *Config) setEnvironments(w *kafka.Writer) {
	if c.MaxAttempts > 0 {
		w.MaxAttempts = c.MaxAttempts
	}
	if c.RequiredAcks > -1 {
		w.RequiredAcks = c.RequiredAcks
	}
	if c.Compression > 0 {
		w.Compression = c.Compression
	}
	if c.WriteTimeOut > 0 {
		w.WriteTimeout = c.WriteTimeOut
	}
	if !c.Async {
		w.Async = c.Async
	}
	if c.ReadTimeOut > 0 {
		w.ReadTimeout = c.ReadTimeOut
	}
	if !c.AllowAutoTopicCreation {
		w.AllowAutoTopicCreation = c.AllowAutoTopicCreation
	}
}
