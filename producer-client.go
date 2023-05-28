package kafka_client

import (
	"google.golang.org/protobuf/proto"
)

type Options func(*BaseKafka)

func WithProtoProducer[T proto.Message](opts ...Options) IProtoProducer[T] {
	p := new(ProtoProducer[T])
	OptionsProcessing(&p.BaseKafka, opts...)
	return p
}

func WithJsonProducer[T any](opts ...Options) IJsonProducer[T] {
	p := new(JsonProducer[T])
	OptionsProcessing(&p.BaseKafka, opts...)
	return p
}

func WithProducer(opts ...Options) IProducer {
	p := new(Producer)
	OptionsProcessing(&p.BaseKafka, opts...)
	return p
}
