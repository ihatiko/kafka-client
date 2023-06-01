package kafka_client

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/segmentio/kafka-go"
)

const (
	topicKey     = "topic"
	offsetKey    = "offset"
	partitionKey = "partition"
)

func GetKafkaTracingHeadersFromSpanCtx(spanCtx opentracing.SpanContext) []kafka.Header {
	textMapCarrier, err := InjectTextMapCarrier(spanCtx)
	if err != nil {
		return []kafka.Header{}
	}

	kafkaMessageHeaders := TextMapCarrierToKafkaMessageHeaders(textMapCarrier)
	return kafkaMessageHeaders
}

func InjectTextMapCarrier(spanCtx opentracing.SpanContext) (opentracing.TextMapCarrier, error) {
	m := make(opentracing.TextMapCarrier)
	if err := opentracing.GlobalTracer().Inject(spanCtx, opentracing.TextMap, m); err != nil {
		return nil, err
	}
	return m, nil
}

func TextMapCarrierToKafkaMessageHeaders(textMap opentracing.TextMapCarrier) []kafka.Header {
	headers := make([]kafka.Header, 0, len(textMap))

	if err := textMap.ForeachKey(func(key, val string) error {
		headers = append(headers, kafka.Header{
			Key:   key,
			Value: []byte(val),
		})
		return nil
	}); err != nil {
		return headers
	}

	return headers
}

func ExtractJaegerContext(message kafka.Message) context.Context {
	m := make(opentracing.TextMapCarrier)
	for _, h := range message.Headers {
		m.Set(h.Key, string(h.Value))
	}
	extractedSpan, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, m)
	if err != nil {
		return context.TODO()
	}
	span := opentracing.StartSpan(
		fmt.Sprintf("Consume.%s.%d", message.Topic, message.Partition),
		opentracing.FollowsFrom(extractedSpan),
	)
	span.SetTag(topicKey, message.Topic)
	span.LogFields(
		log.Int64(offsetKey, message.Offset),
		log.Int(partitionKey, message.Partition),
	)

	return opentracing.ContextWithSpan(context.TODO(), span)
}
