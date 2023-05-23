package kafka_client

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

type Producer struct {
	cfg    *TopicConfig
	conCfg *Config
	w      *kafka.Writer
}

func (c *Config) NewProducer(ctx context.Context, topicConfig *TopicConfig) (*Producer, error) {
	err := c.checkConnection(ctx, topicConfig)
	if err != nil {
		return nil, err
	}
	return &Producer{
		cfg:    topicConfig,
		w:      c.NewWriter(),
		conCfg: c,
	}, err
}

func (p *Producer) Publish(ctx context.Context, data ...[]byte) error {
	for _, d := range data {
		message := kafka.Message{
			Value: d,
			Topic: p.cfg.Name,
			Time:  time.Now(),
		}
		err := p.w.WriteMessages(ctx, message)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *Producer) PublishJson(ctx context.Context, data ...[]any) error {
	for _, d := range data {
		m, err := jsoniter.Marshal(d)
		if err != nil {
			return err
		}
		err = p.publish(ctx, m)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *Producer) PublishProto(ctx context.Context, data ...proto.Message) error {
	for _, d := range data {
		m, err := proto.Marshal(d)
		if err != nil {
			return err
		}
		err = p.publish(ctx, m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Producer) publish(ctx context.Context, data []byte) error {
	span, ctxSpan := opentracing.StartSpanFromContext(
		ctx,
		fmt.Sprintf("Publish_To_Topic.%s", p.cfg.Name),
	)
	message := kafka.Message{
		Value:   data,
		Topic:   p.cfg.Name,
		Time:    time.Now(),
		Headers: GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}
	return p.w.WriteMessages(ctxSpan, message)
}
