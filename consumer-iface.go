package kafka_client

import (
	"context"
)

type IConsumer[T any] interface {
	Consume(context.Context, T) error
}
