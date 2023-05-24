package kafka_client

import (
	"context"
	"google.golang.org/protobuf/proto"
)

type IConsumer interface {
	Consume(context.Context, []byte) error
}

type IProtoConsumer[T proto.Message] interface {
	Consume(context.Context, T) error
}

type IJsonConsumer[T any] interface {
	Consume(context.Context, T) error
}
