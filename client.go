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

type Queue struct {
	Error error
}

func Health(hosts ...string) error {
	var (
		globalErr error
	)
	wg := sync.WaitGroup{}
	for _, br := range hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			var curAttempts int
			for {
				_, err := kafka.DialContext(context.TODO(), "tcp", host)
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
	return globalErr
}
