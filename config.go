package kafka_client

import (
	"github.com/segmentio/kafka-go"
	"time"
)

type TopicConfig struct {
	Name              string
	ReplicationFactor int
	Partitions        int
}

type ConsumerGroup struct {
	GroupID string
	Topics  []string
	DLQ     *Backoff
	Debug   bool
}

type Config struct {
	Host                   []string
	WriteTimeOut           time.Duration
	ReadTimeOut            time.Duration
	AsyncDefault           bool
	DialTimeout            time.Duration
	MaxAttempts            int
	PartitionWatchInterval time.Duration
	CommitIntervalDefault  int
	HeartbeatInterval      time.Duration
	QueueCapacity          int
	MinBytes               float64 // 10KB
	MaxBytes               float64 // 10MB
	Compression            kafka.Compression
	Async                  bool
	UseSSL                 bool
	SslCaPem               string
	Username               string
	Password               string
	AllowAutoTopicCreation bool
	RequiredAcks           kafka.RequiredAcks
}
