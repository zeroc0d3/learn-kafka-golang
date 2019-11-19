package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/zeroc0d3/learn-kafka-golang/src/producer"
)

var (
	// Word list dummy messages (read from file)
	wordList *os.File
	err      error
)

func main() {
	// Setup Logging
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := getKafkaConfig("", "")
	producers, err := sarama.NewSyncProducer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return
	}
	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	logrus.Infof("Success create kafka sync-producer")

	kafka := &producer.KafkaProducer{
		Producer: producers,
	}

	wordList, err := os.OpenFile("words.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}

	scanner := bufio.NewScanner(wordList)
	scanner.Split(bufio.ScanLines)

	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	defer wordList.Close()

	// Get topic name.
	topicName := os.Getenv("KAFKA_TOPIC")

	// Test With Number
	// for i := 1; i <= 10; i++ {
	// 	msg := fmt.Sprintf("message number %v", i)
	// 	err := kafka.SendMessage(topicName, msg)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// Using Word List
	for _, eachline := range txtlines {
		result := kafka.SendMessage(topicName, eachline)

		// Block until the result is returned and a server-generated
		// ID is returned for the published message.
		id := 1
		err := result
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("## Send Message [ %v => %s ].\n", id, eachline)
		id = id + 1

		// Test show wordlist
		// fmt.Println(eachline)
	}

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
