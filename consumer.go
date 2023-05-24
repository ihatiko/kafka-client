package kafka_client

import (
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type ConsumerGroup struct {
	Brokers []string
	GroupID string
	log     kafka.Logger
	cfgConn *Config
	DLQ     bool
}

func ConsumeTopic(consumer IConsumer) {

}

func ConsumeProtoTopic[T proto.Message](consumer IProtoConsumer[T]) {

}

func ConsumeJsonTopic[T any](consumer IJsonConsumer[T]) {

}
