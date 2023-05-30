package kafka_client

import "golang.org/x/net/context"

type Handler[T any] func(request *Request[T]) error

type Data []byte

type Request[T any] struct {
	Context context.Context
	Data    T
}
