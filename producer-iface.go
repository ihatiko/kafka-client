package kafka_client

import (
	"context"
	"google.golang.org/protobuf/proto"
)

type IProducer interface {
	Publish(ctx context.Context, data ...[]byte) error
	PublishJson(ctx context.Context, data ...[]any) error
	PublishProto(ctx context.Context, data ...proto.Message) error
}
