package main
import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func subscribe(topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetNewest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		log.Printf("topic: [%s]", topic)
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				messageReceived(message)
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {
	log.Printf("msg : [%s]", message.Value)
	saveMessage(string(message.Value))
}
