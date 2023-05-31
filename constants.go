package kafka_client

import (
	"time"

	"github.com/segmentio/kafka-go/compress"
)

const (
	WriterRequiredAcks = -1
)

const (
	WriteTimeOut           = 15 * time.Second
	ReadTimeOut            = 15 * time.Second
	Compression            = compress.Gzip
	AsyncDefault           = false
	DialTimeout            = 3 * time.Minute
	MaxAttempts            = 10
	PartitionWatchInterval = 5 * time.Second
	CommitIntervalDefault  = 0
	HeartbeatInterval      = 3 * time.Second
	QueueCapacity          = 100
	MinBytes               = 10e3 // 10KB
	MaxBytes               = 10e6 // 10MB
	AllowAutoTopicCreation = true
)
