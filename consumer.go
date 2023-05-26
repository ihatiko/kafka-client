package kafka_client

import (
	"google.golang.org/protobuf/proto"
)

func (c *Config) NewConsumer(cGroup *ConsumerGroup) {
	//c.checkConnection(context.Background())
	//c.newReader(cGroup.GroupID, cGroup.Topics...)
}

func ConsumeTopic(consumer IConsumer) {

}

func ConsumeProtoTopic[T proto.Message](consumer IProtoConsumer[T]) {

}

func ConsumeJsonTopic[T any](consumer IJsonConsumer[T]) {

}
