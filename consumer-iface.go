package kafka_client

import (
	"google.golang.org/protobuf/proto"
)

type IConsumer interface {
	Consume(handler Handler[Data])
}

type IProtoConsumer[T proto.Message] interface {
	Consume(handler Handler[T])
}

type IJsonConsumer[T any] interface {
	Consume(handler Handler[T])
}
