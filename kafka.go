package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

var (
	brokers = []string{"127.0.0.1:9092"}
	topic   = "xbanku-transactions-t3"
	topics  = []string{topic}
)

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0
	return conf
}

func newKafkaSyncProducer() sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("Kafka sync-producer error: %s\n", err)
		os.Exit(-1)
	}
	return producer
}

func newKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("Kafka consumer error: %s\n", err)
	}
	return consumer
}

func sendMsg(producer sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msgLog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := producer.SendMessage(msgLog)
	if err != nil {
		fmt.Printf("Kafka error in sending a message %s\n", err)
	}

	fmt.Printf("Message %v\n", event)
	fmt.Printf("Message is stored in partition %d, offset %d\n", partition, offset)

	return nil
}
