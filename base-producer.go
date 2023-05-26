package kafka_client

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"time"
)

type BaseProducer struct {
	err          error
	TopicConfig  *TopicConfig
	WriterConfig *Config
	Writer       *kafka.Writer
}

func (t *BaseProducer) Error() error {
	return t.err
}

func (t *BaseProducer) HasError() bool {
	return t.err != nil
}

func (t *BaseProducer) publish(ctx context.Context, data []byte) error {
	span, ctxSpan := opentracing.StartSpanFromContext(
		ctx,
		fmt.Sprintf("Publish.%s", t.TopicConfig.Name),
	)
	message := kafka.Message{
		Value:   data,
		Topic:   t.TopicConfig.Name,
		Time:    time.Now(),
		Headers: GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}
	return t.Writer.WriteMessages(ctxSpan, message)
}
