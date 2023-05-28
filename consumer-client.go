package kafka_client

func WithConsumer(opt ...Options) IConsumer {
	consumer := new(Consumer)
	OptionsProcessing(&consumer.BaseKafka, opt...)
	return consumer
}
