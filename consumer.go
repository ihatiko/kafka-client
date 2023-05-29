package kafka_client

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	attemptKey = "attempt"
	waitForKey = "waitFor"
)

type Consumer struct {
	BaseKafka
}

func (c *Consumer) Consume(handler Handler[Data]) {
	ctx := context.Background()
	for _, rd := range c.Readers {
		go c.consume(ctx, rd, handler)
	}
}

func (c *Consumer) consume(ctx context.Context, reader *kafka.Reader, handler Handler[Data]) {
	//TODO shutdown
	fmt.Println(fmt.Sprintf("start consume %v", reader.Config().GroupTopics))
	for {
		m, err := reader.FetchMessage(ctx)
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
			deliveryTopic := m.Topic
			attemptHeader, ok := findHeader(attemptKey, m.Headers)
			var waitForHeader *kafka.Header
			if !ok {
				attemptHeader = &kafka.Header{
					Key:   attemptKey,
					Value: []byte("1"),
				}
				nextDur := c.ConsumerConfig.DLQ.GetDuration(1)
				nextTime := time.Now().Add(nextDur)
				valueHeaderDate, _ := nextTime.GobEncode()
				waitForHeader = &kafka.Header{
					Key:   waitForKey,
					Value: valueHeaderDate,
				}
				deliveryTopic = fmt.Sprintf("%s.attemtps.%d", m.Topic, 1)
			} else {
				currentAttempt, _ := strconv.Atoi(string(attemptHeader.Value))
				currentAttempt += 1
				if currentAttempt == c.ConsumerConfig.DLQ.Attempts {
					deliveryTopic = fmt.Sprintf("%s.DLQ", deliveryTopic)
				} else {
					nextDur := c.ConsumerConfig.DLQ.GetDuration(currentAttempt)
					nextTime := time.Now().Add(nextDur)
					valueHeaderDate, _ := nextTime.GobEncode()
					waitForHeader = &kafka.Header{
						Key:   waitForKey,
						Value: valueHeaderDate,
					}
					attemptHeader.Value = []byte(strconv.Itoa(currentAttempt))
					currentMainTopic := strings.Split(m.Topic, ".attempts")[0]
					deliveryTopic = fmt.Sprintf("%s.attemtps.%d", currentMainTopic, currentAttempt)
				}
			}
			if extractedContext == nil {
				extractedContext = context.TODO()
			}
			err = c.publish(
				extractedContext,
				deliveryTopic,
				m.Value,
				*attemptHeader,
				*waitForHeader,
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

func findHeader(key string, headers []kafka.Header) (*kafka.Header, bool) {
	for _, h := range headers {
		if h.Key == key {
			return &h, true
		}
	}
	return nil, false
}
