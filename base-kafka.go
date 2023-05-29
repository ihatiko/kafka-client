package kafka_client

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"time"
)

type BaseKafka struct {
	err            error
	TopicConfig    *TopicConfig
	ConsumerConfig *ConsumerGroup
	DLQ            map[string]*kafka.Writer
	Readers        map[string]*kafka.Reader
	KafkaConfig    *Config
	Writer         *kafka.Writer
}

func (t *BaseKafka) Error() error {
	return t.err
}

func (t *BaseKafka) HasError() bool {
	return t.err != nil
}

func (t *BaseKafka) publish(ctx context.Context, topic string, data []byte, headers ...kafka.Header) error {
	span, ctxSpan := opentracing.StartSpanFromContext(
		ctx,
		fmt.Sprintf("Publish.%s", topic),
	)
	headers = append(headers, GetKafkaTracingHeadersFromSpanCtx(span.Context())...)
	message := kafka.Message{
		Value:   data,
		Topic:   topic,
		Time:    time.Now(),
		Headers: headers,
	}
	return t.Writer.WriteMessages(ctxSpan, message)
}
