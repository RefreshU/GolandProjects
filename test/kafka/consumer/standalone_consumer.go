package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"time"
)

// SinglePartition 单分区消费
func SinglePartition(topic string) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1"}, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}

	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] PartitionId: %d; offset: %d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}

// Partitions 多分区消费
func Partitions(topic string){
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1"}, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	defer consumer.Close()
	// 先查询该topic 有多少分区
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatal("Partitins err: ", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(partitions))
	for _, partitionId := range partitions	{
		go consumeByPartition(consumer, topic, partitionId, &wg)
	}

}

func consumeByPartition(consumer sarama.Consumer, topic string, partitionId int32, wg *sync.WaitGroup){
	defer wg.Done()
	partitionConsumer, err := consumer.ConsumePartition(topic, partitionId, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consume] PartitionId: %d; offset: %d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}

func OffsetManager(topic string){
	config := sarama.NewConfig()

	config.Consumer.Offsets.AutoCommit.Enable = true // 开启自动 commit offset
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second //自动commit时间间隔
	client, err := sarama.NewClient([]string{"127.0.0.1"}, config)
	if err != nil {
		log.Fatal("NewClient err: ", err)
	}
	defer client.Close()

	offsetManager, err := sarama.NewOffsetManagerFromClient("GroupId", client)
	if err != nil {
		log.Fatal("NewOffsetManagerFromClient err: ", err)
	}
	defer offsetManager.Close()
	//每个分区的offset分别管理
	partitionOffsetManager, err := offsetManager.ManagePartition(topic, 0)
	if err != nil {
		log.Fatal("ManagePartition err: ", err)
	}
	defer partitionOffsetManager.Close()
	// defer 在程序结束后再commit 一次， 防止自动提交间隔之间的信息被丢掉
	defer offsetManager.Commit()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatal("NewConsumerFromClient err: ", err)
	}
	// 取得下一个消息的偏移量作为本次消费的起点
	nextOffset, _ := partitionOffsetManager.NextOffset()
	fmt.Println("nextOffset: ", nextOffset)
	pc, err := consumer.ConsumePartition(topic, 0, nextOffset)
	if err != nil {
		log.Println("ConsumePartition err: ", err)
	}
	defer pc.Close()

	for message := range pc.Messages() {
		value := string(message.Value)
		log.Println("[Consumer] partitionId: %d; offset: %d; value: %s\n", message.Partition, message.Offset, value)
		partitionOffsetManager.MarkOffset(message.Offset+1, "modified metadata")
	}
}