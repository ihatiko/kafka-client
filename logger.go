package kafka_client

type ILogger interface {
	WarnF()
	InfoF()
	FatalF()
}
