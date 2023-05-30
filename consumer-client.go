package kafka_client

import "google.golang.org/protobuf/proto"

func WithConsumer(opt ...Options) IConsumer {
	consumer := new(Consumer)
	OptionsProcessing(&consumer.BaseKafka, opt...)
	return consumer
}

func WithJsonConsumer[T any](opt ...Options) IJsonConsumer[T] {
	consumer := new(ConsumerJson[T])
	OptionsProcessing(&consumer.BaseKafka, opt...)
	return consumer
}

func WithProtoConsumer[T proto.Message](opt ...Options) IProtoConsumer[T] {
	consumer := new(ConsumerProto[T])
	OptionsProcessing(&consumer.BaseKafka, opt...)
	return consumer
}
