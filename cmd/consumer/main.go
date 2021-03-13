package main

import (
	"os"
	"time"

	"github.com/zeroc0d3/learn-kafka-golang/src/consumer"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func main() {
	// Setup Logging
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := getKafkaConfig("", "")
    envBrokerConfig := os.Getenv("KAFKA_BROKER_URL")

	// consumers, err := sarama.NewConsumer([]string{"kafka:9092"}, kafkaConfig)
    consumers, err := sarama.NewConsumer([]string{ envBrokerConfig }, kafkaConfig)
	if err != nil {
		logrus.Errorf("Error create kakfa consumer got error %v", err)
	}
	defer func() {
		if err := consumers.Close(); err != nil {
			logrus.Fatal(err)
			return
		}
	}()

	kafkaConsumer := &consumer.KafkaConsumer{
		Consumer: consumers,
	}

	signals := make(chan os.Signal, 1)

	// Get topic name.
	topicName := os.Getenv("KAFKA_TOPIC")
	kafkaConsumer.Consume([]string{topicName}, signals)
}

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
