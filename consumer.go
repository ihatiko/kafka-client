package kafka_client

import (
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

type ConsumerGroup struct {
	Brokers []string
	GroupID string
	log     kafka.Logger
	cfgConn *Config
}

type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

func (c *ConsumerGroup) ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker) {
	r := c.cfgConn.newReader(c.GroupID, groupTopics...)

	defer func() {
		if err := r.Close(); err != nil {
			c.log.Printf("consumerGroup.r.Close: %v", err)
		}
	}()

	c.log.Printf("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)
	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize+1; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}
