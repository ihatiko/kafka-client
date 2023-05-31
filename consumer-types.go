package kafka_client

import "golang.org/x/net/context"

type MetaKey string

type Handler[T any] func(request *Request[T]) error

type Data []byte

type Request[T any] struct {
	Context context.Context
	Data    T
	Meta    map[MetaKey]string
}
