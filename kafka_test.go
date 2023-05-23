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

func Test_kafka(t *testing.T) {
	//set := dockertest.WithPool().
	//	WithZookeeper().
	//	WithKafka()
	//assert.NilError(t, set.Error)

	cfg := Config{
		InitTopics: true,
		Host:       []string{"localhost:60001"},
	}

	writer, err := cfg.NewProducer(context.Background(), &TopicConfig{
		Name:              "sandbox",
		Partitions:        1,
		ReplicationFactor: 1,
	})

	assert.NilError(t, err)
	err = writer.Publish(context.TODO())
	assert.NilError(t, err)
}
