package kafka_client

import (
	"fmt"
	"github.com/golang/protobuf/descriptor"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func WithConfig(cfg *Config) Options {
	return func(t *BaseKafka) {
		t.KafkaConfig = cfg
	}
}

func WithConsumerGroup(cfg *ConsumerGroup) Options {
	return func(t *BaseKafka) {
		t.ConsumerConfig = cfg
		t.Readers = make(map[string]*kafka.Reader)
		for _, tp := range cfg.Topics {

			t.Readers[tp] = t.KafkaConfig.newReader(t.ConsumerConfig)
		}
		t.Writer = t.KafkaConfig.NewWriter()
	}
}

func WithDLQConsumerGroup(cfg *ConsumerGroup) Options {
	return func(t *BaseKafka) {
		t.ConsumerConfig = cfg
		t.Writer = t.KafkaConfig.NewWriter()
		t.Readers = make(map[string]*kafka.Reader)
		for _, topic := range cfg.Topics {
			for i := 1; i < cfg.DLQ.Attempts+1; i++ {
				key := fmt.Sprintf("%s.attempts.%d", topic, i)
				t.Readers[key] = t.KafkaConfig.newReader(&ConsumerGroup{
					Topics:  []string{key},
					GroupID: cfg.GroupID,
				})
			}
		}
	}
}

func WithHealth(hosts []string) Options {
	return func(t *BaseKafka) {
		t.err = Health(hosts...)
	}
}

func WithTopic(cfg *TopicConfig) Options {
	return func(t *BaseKafka) {
		t.TopicConfig = cfg
		t.Writer = t.KafkaConfig.NewWriter()
	}
}

func WithProtoTopic[T any](xt protoreflect.ExtensionType) Options {
	return func(t *BaseKafka) {
		_, md := descriptor.MessageDescriptorProto(new(T))
		ex := proto.GetExtension(md.Options, xt)
		t.TopicConfig = &TopicConfig{
			Name: ex.(string),
		}
		t.Writer = t.KafkaConfig.NewWriter()
	}
}
