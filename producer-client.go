package kafka_client

import (
	"github.com/golang/protobuf/descriptor"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ProducerOptions func(*BaseProducer)

func WithHealth(hosts []string) ProducerOptions {
	return func(t *BaseProducer) {
		t.err = Health(hosts...)
	}
}

func WithConfig(cfg *Config) ProducerOptions {
	return func(t *BaseProducer) {
		t.WriterConfig = cfg
		t.Writer = cfg.NewWriter()
	}
}

func WithTopic(topicCfg *TopicConfig) ProducerOptions {
	return func(t *BaseProducer) {
		t.TopicConfig = topicCfg
	}
}

func WithProtoTopic[T any](xt protoreflect.ExtensionType) ProducerOptions {
	return func(t *BaseProducer) {
		_, md := descriptor.MessageDescriptorProto(new(T))
		ex := proto.GetExtension(md.Options, xt)
		t.TopicConfig = &TopicConfig{
			Name: ex.(string),
		}
	}
}

func OptionsProcessing(data *BaseProducer, opts ...ProducerOptions) {
	for _, opt := range opts {
		if data.HasError() {
			continue
		}
		opt(data)
	}
}
func WithProtoProducer[T proto.Message](opts ...ProducerOptions) IProtoProducer[T] {
	p := new(ProtoProducer[T])
	OptionsProcessing(&p.BaseProducer, opts...)
	return p
}

func WithJsonProducer[T any](opts ...ProducerOptions) IJsonProducer[T] {
	p := new(JsonProducer[T])
	OptionsProcessing(&p.BaseProducer, opts...)
	return p
}

func WithProducer(opts ...ProducerOptions) IProducer {
	p := new(Producer)
	OptionsProcessing(&p.BaseProducer, opts...)
	return p
}
