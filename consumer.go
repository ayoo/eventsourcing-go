package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func mainConsumer(partition int32) {
	consumer := newKafkaConsumer()
	defer consumer.Close()

	pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("Kafka error: %s", err)
		os.Exit(-1)
	}

	go consumeEvents(pc)

	fmt.Println("Exiting mainConsumer")
	fmt.Println("Press [enter] to exit consumer\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Terminating...")

}

func consumeEvents(pc sarama.PartitionConsumer) {
	var msgVal []byte
	var log interface{}
	var logMap map[string]interface{}

	fmt.Println("Waiting a message from the channel")
	for {
		select {
		case err := <-pc.Errors():
			fmt.Printf("Kafka consumer error: %s\n", err)
		case msg := <-pc.Messages():
			msgVal = msg.Value
			err := json.Unmarshal(msgVal, &log)
			if err != nil {
				fmt.Printf("Failed Parsing: %s", err)
				return
			}
			logMap = log.(map[string]interface{})
			logType := logMap["Type"]

			fmt.Printf("Processing %s:\n%s\n", logMap["Type"], string(msgVal))
			handleEvent(logType, msgVal)
		}
	}
}

func handleEvent(logType interface{}, msgVal []byte) {
	var err error
	var bankAccount *BankAccount

	switch logType {
	case "CreateEvent":
		bankAccount, err = processEvent(new(CreateEvent), msgVal)
	case "DepositEvent":
		bankAccount, err = processEvent(new(DepositEvent), msgVal)
	case "WithdrawEvent":
		bankAccount, err = processEvent(new(WithdrawEvent), msgVal)
	case "TransferEvent":
		event := new(TransferEvent)
		bankAccount, err = processEvent(event, msgVal)
		if err == nil {
			if targetAcc, err := GetAccount(event.TargetID); err == nil {
				fmt.Printf("%+v\n", *targetAcc)
			}
		}
	default:
		fmt.Println("Unknown command: ", logType)
	}

	if err != nil {
		fmt.Printf("Error processing: %s\n", err)
		return
	}
	fmt.Printf("%+v\n\n", *bankAccount)
}

func processEvent(event EventAction, msgVal []byte) (*BankAccount, error) {
	err := json.Unmarshal(msgVal, &event)
	if err != nil {
		return nil, err
	}
	return event.Process()
}
