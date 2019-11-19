package consumer

import (
	"os"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// KafkaConsumer hold sarama consumer
type KafkaConsumer struct {
	Consumer sarama.Consumer
}

// Consume function to consume message from apache kafka
func (c *KafkaConsumer) Consume(topics []string, signals chan os.Signal) {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)

	for _, topic := range topics {
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("Unable to get partition got error %v", err)
			continue
		}
		for _, partition := range partitionList {
			go consumeMessage(c.Consumer, topic, partition, chanMessage)
		}
	}
	logrus.Infof("Kafka is consuming....")

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			logrus.Infof("New Message from kafka, message: %v", string(msg.Value))
		case sig := <-signals:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		logrus.Errorf("Unable to consume partition %v got error %v", partition, err)
		return
	}

	defer func() {
		if err := msg.Close(); err != nil {
			logrus.Errorf("Unable to close partition %v: %v", partition, err)
		}
	}()

	for {
		msg := <-msg.Messages()
		c <- msg
	}

}
