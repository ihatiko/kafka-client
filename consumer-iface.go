package kafka_client

import (
	"google.golang.org/protobuf/proto"
)

type IConsumer interface {
	ErrorComposition
	Consume(handler Handler[Data])
}

type IProtoConsumer[T proto.Message] interface {
	ErrorComposition
	Consume(handler Handler[T])
}

type IJsonConsumer[T any] interface {
	ErrorComposition
	Consume(handler Handler[T])
}
