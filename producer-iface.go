package kafka_client

import (
	"context"
	"google.golang.org/protobuf/proto"
)

type IProducer interface {
	ErrorComposition
	Publish(ctx context.Context, data ...[]byte) error
}

type IProtoProducer[T proto.Message] interface {
	ErrorComposition
	Publish(ctx context.Context, data ...T) error
}

type IJsonProducer[T any] interface {
	ErrorComposition
	Publish(ctx context.Context, data ...T) error
}
