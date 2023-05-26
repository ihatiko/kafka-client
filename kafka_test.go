package kafka_client

import (
	"context"
	"gotest.tools/v3/assert"
	"kafka-client/protoc/example"
	"testing"
)

type TestData struct {
	Id   string
	Name string
}

func Test_kafka_producer(t *testing.T) {
	//set := dockertest.WithPool().
	//	WithZookeeper().
	//	WithKafka()
	//assert.NilError(t, set.Error)

	cfg := &Config{
		InitTopics: true,
		Host:       []string{"localhost:9092"},
	}

	tpCfg := &TopicConfig{
		Name:              "sandbox",
		Partitions:        1,
		ReplicationFactor: 1,
	}

	writer := WithProducer(
		WithHealth(cfg.Host),
		WithConfig(cfg),
		WithTopic(tpCfg),
	)

	assert.NilError(t, writer.Error())
	err := writer.Publish(context.TODO(), []byte(`{"name": "tests"}`))
	assert.NilError(t, err)

	writer2 := WithProtoProducer[*example.ExampleMessageRequest](
		WithHealth(cfg.Host),
		WithConfig(cfg),
		WithProtoTopic[example.ExampleMessageRequest](example.E_Topic),
	)

	err = writer2.Publish(context.TODO(), &example.ExampleMessageRequest{
		Name1: "test1",
		Name2: "test2",
		Name3: "test3",
		Name4: "test4",
		Name5: "test5",
	})

	assert.NilError(t, err)

	writer3 := WithJsonProducer[*TestData](
		WithHealth(cfg.Host),
		WithConfig(cfg),
		WithTopic(tpCfg),
	)

	err = writer3.Publish(context.TODO(), &TestData{
		Name: "test",
		Id:   "123",
	})

	assert.NilError(t, err)
}

func Test_kafka_consumers(t *testing.T) {
	//set := dockertest.WithPool().
	//	WithZookeeper().
	//	WithKafka()
	//assert.NilError(t, set.Error)

	//cfg := Config{
	//	InitTopics: true,
	//	Host:       []string{"localhost:9092"},
	//}
	//
	//writer, err := cfg.NewProducer(context.Background(), &TopicConfig{
	//	Name:              "sandbox",
	//	Partitions:        1,
	//	ReplicationFactor: 1,
	//})
	//
	//assert.NilError(t, err)
	//err = writer.Publish(context.TODO(), []byte(`{"name": "tests"}`))
	//assert.NilError(t, err)
}
