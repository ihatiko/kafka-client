package kafka_client

import (
	"time"

	"github.com/segmentio/kafka-go/compress"
)

const (
	minBytesDefault               = 10e3 // 10KB
	maxBytesDefault               = 10e6 // 10MB
	queueCapacityDefault          = 100
	heartbeatIntervalDefault      = 3 * time.Second
	commitIntervalDefault         = 0
	partitionWatchIntervalDefault = 5 * time.Second
	maxAttemptsDefault            = 3
	dialTimeoutDefault            = 3 * time.Minute
	maxWaitDefault                = 1 * time.Second

	WriterReadTimeoutDefault  = 10 * time.Second
	WriterWriteTimeoutDefault = 10 * time.Second
	WriterRequiredAcksDefault = -1
	CompressionDefault        = compress.Gzip
	AsyncDefault              = false
)
