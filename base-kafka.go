package kafka_client

import (
	"context"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
	"strings"
	"time"
)

type BaseKafka struct {
	err            error
	TopicConfig    *TopicConfig
	ConsumerConfig *ConsumerGroup
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
	err := retry.Do(func() error {
		return t.Writer.WriteMessages(ctxSpan, message)
	},
		//TODO cfg
		retry.MaxDelay(t.KafkaConfig.MaxWait*time.Second),
		retry.LastErrorOnly(true),
		retry.Delay(time.Millisecond*100),
	)

	return err
}

func (t *BaseKafka) consume(ctx context.Context, reader *kafka.Reader, handler func(context.Context, []byte) error) {
	//TODO shutdown
	fmt.Println(fmt.Sprintf("start consume %v", reader.Config().GroupTopics))
	for {
		var (
			headers []kafka.Header
		)
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			break
		}
		extractedContext := ExtractJaegerContext(m)
		err = handler(extractedContext, m.Value)
		if h, ok := findHeader(waitForKey, m.Headers); ok {
			potentialTime := time.Time{}
			err := potentialTime.GobDecode(h.Value)
			if err != nil {
				log.Fatalf("Broken date format topic: %s offset: %d", m.Topic, m.Offset)
			}
			if potentialTime.After(time.Now()) {
				dif := potentialTime.Sub(time.Now())
				time.Sleep(dif)
			}
		}
		if err != nil && t.ConsumerConfig.DLQ != nil && !strings.HasSuffix(m.Topic, DLQKey) {
			deliveryTopic := m.Topic
			attemptHeader, ok := findHeader(attemptKey, m.Headers)
			if !ok {
				attemptHeader := kafka.Header{
					Key:   attemptKey,
					Value: []byte("1"),
				}
				nextDur := t.ConsumerConfig.DLQ.GetDuration(1)
				nextTime := time.Now().Add(nextDur)
				valueHeaderDate, _ := nextTime.GobEncode()
				waitForHeader := kafka.Header{
					Key:   waitForKey,
					Value: valueHeaderDate,
				}
				deliveryTopic = fmt.Sprintf("%s.%s.%d", m.Topic, attemptKey, 1)
				headers = append(headers, waitForHeader, attemptHeader)
			} else {
				currentAttempt, _ := strconv.Atoi(string(attemptHeader.Value))
				currentAttempt += 1
				if currentAttempt > t.ConsumerConfig.DLQ.Attempts {
					currentMainTopic := strings.Split(m.Topic, replaceAttemptKey)[0]
					deliveryTopic = fmt.Sprintf("%s.%s", currentMainTopic, DLQKey)
				} else {
					nextDur := t.ConsumerConfig.DLQ.GetDuration(currentAttempt)
					nextTime := time.Now().Add(nextDur)
					valueHeaderDate, _ := nextTime.GobEncode()
					waitForHeader := kafka.Header{
						Key:   waitForKey,
						Value: valueHeaderDate,
					}
					attemptHeader.Value = []byte(strconv.Itoa(currentAttempt))
					currentMainTopic := strings.Split(m.Topic, replaceAttemptKey)[0]
					deliveryTopic = fmt.Sprintf("%s.%s.%d", currentMainTopic, attemptKey, currentAttempt)
					headers = append(headers, waitForHeader, attemptHeader)
				}
			}
			err = t.publish(
				extractedContext,
				deliveryTopic,
				m.Value,
				headers...,
			)
			if err != nil {
				log.Fatal("failed to commit dlq message:", err)
			}
		}
		if err := reader.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}

func findHeader(key string, headers []kafka.Header) (kafka.Header, bool) {
	for _, h := range headers {
		if h.Key == key {
			return h, true
		}
	}
	return kafka.Header{}, false
}
