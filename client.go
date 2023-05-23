package kafka_client

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"sync"
	"time"
)

const (
	attempts = 10
	delay    = 1
)

func (c *Config) checkConnection(ctx context.Context, topicConfig *TopicConfig) error {
	var (
		conn      *kafka.Conn
		err       error
		globalErr error
	)
	wg := sync.WaitGroup{}
	for _, br := range c.Host {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			var curAttempts int
			for {
				conn, err = kafka.DialContext(ctx, "tcp", host)
				if err == nil {
					break
				}
				if err != nil {
					curAttempts++
				}
				if attempts < curAttempts {
					err = errors.Join(globalErr, err)
					break
				}
				time.Sleep(delay * time.Second)
			}
		}(br)
	}
	wg.Wait()
	if globalErr != nil {
		return globalErr
	}
	if c.InitTopics {
		err = conn.CreateTopics(kafka.TopicConfig{
			Topic:             topicConfig.Name,
			ReplicationFactor: topicConfig.ReplicationFactor,
			NumPartitions:     topicConfig.Partitions,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
