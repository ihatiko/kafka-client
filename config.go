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
	DLQ     bool
}

type Config struct {
	InitTopics             bool
	Host                   []string
	MinBytes               int
	MaxBytes               int
	QueueCapacity          int
	HeartbeatInterval      time.Duration
	CommitInterval         int
	PartitionWatchInterval time.Duration
	MaxAttempts            int
	DialTimeout            time.Duration
	MaxWait                time.Duration
	WriterReadTimeout      time.Duration
	WriterWriteTimeout     time.Duration
	WriterRequiredAcks     int
	WriterMaxAttempts      int
	Compression            kafka.Compression
	Async                  bool
	UseSSL                 bool
	SslCaPem               string
	Username               string
	Password               string
}
