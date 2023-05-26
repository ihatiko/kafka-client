package kafka_client

type ErrorComposition interface {
	HasError() bool
	Error() error
}
