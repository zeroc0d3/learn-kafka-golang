package consumer_test

import (
	"os"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	"github.com/zeroc0d3/learn-kafka-golang/src/consumer"
)

func TestConsume(t *testing.T) {
	consumers := mocks.NewConsumer(t, nil)
	defer func() {
		if err := consumers.Close(); err != nil {
			t.Error(err)
		}
	}()

	// Get topic name.
	topicName := os.Getenv("KAFKA_TOPIC")

	consumers.SetTopicMetadata(map[string][]int32{
		topicName: {0},
	})

	kafka := &consumer.KafkaConsumer{
		Consumer: consumers,
	}

	consumers.ExpectConsumePartition(topicName, 0, sarama.OffsetNewest).YieldMessage(&sarama.ConsumerMessage{Value: []byte("hello world")})

	signals := make(chan os.Signal, 1)
	go kafka.Consume([]string{topicName}, signals)
	timeout := time.After(2 * time.Second)
	for {
		select {
		case <-timeout:
			signals <- os.Interrupt
			return
		}
	}
}
