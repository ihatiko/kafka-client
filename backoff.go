package kafka_client

import (
	"math"
	"time"
)

type Backoff struct {
	Attempts int
	Factor   float64
	MaxDelay time.Duration
}

func (b *Backoff) GetDuration(currentAttempt int) time.Duration {
	d := time.Duration(math.Pow(b.Factor, float64(currentAttempt))) * time.Second
	if d > b.MaxDelay*time.Minute {
		return b.MaxDelay * time.Minute
	}
	return d
}
