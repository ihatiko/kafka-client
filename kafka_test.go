package kafka_client

import (
	"context"
	"gotest.tools/v3/assert"
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

	cfg := Config{
		InitTopics: true,
		Host:       []string{"localhost:9092"},
	}

	writer, err := cfg.NewProducer(context.Background(), &TopicConfig{
		Name:              "sandbox",
		Partitions:        1,
		ReplicationFactor: 1,
	})

	assert.NilError(t, err)
	err = writer.Publish(context.TODO(), []byte(`{"name": "tests"}`))
	assert.NilError(t, err)
}

func Test_kafka_consumers(t *testing.T) {
	//set := dockertest.WithPool().
	//	WithZookeeper().
	//	WithKafka()
	//assert.NilError(t, set.Error)

	cfg := Config{
		InitTopics: true,
		Host:       []string{"localhost:9092"},
	}

	writer, err := cfg.NewProducer(context.Background(), &TopicConfig{
		Name:              "sandbox",
		Partitions:        1,
		ReplicationFactor: 1,
	})

	assert.NilError(t, err)
	err = writer.Publish(context.TODO(), []byte(`{"name": "tests"}`))
	assert.NilError(t, err)
}
