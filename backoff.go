package kafka_client

import (
	"math"
	"math/rand"
	"time"
)

type Backoff struct {
	Attempts    int
	MaxAttempts int
	Factor      float64
	Max         time.Duration
}

func (b *Backoff) NextAttempt() time.Duration {
	return GetNextDuration(1, b.Max, b.Factor, b.Attempts)
}

func GetNextDuration(min, max time.Duration, factor float64, attempts int) time.Duration {
	d := time.Duration(float64(min) * math.Pow(factor, float64(attempts)))
	if d > max*time.Minute {
		return max * time.Minute
	}
	return d
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
