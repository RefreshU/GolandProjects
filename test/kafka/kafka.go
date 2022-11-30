// The Kafka documentation makes it very confusing to set up plain text SASL authentication while also using TLS / SSL.
// MAKE SURE THE KEYSTORE YOU ARE USING ON THE KAFKA CLUSTER IS BUILT WITH RSA ALGO, OTHERWISE GO CAN'T TALK TO JAVA OVER TLS / SSL
package main

import (
	"crypto/tls"
	"fmt"
	"github.com/Shopify/sarama"
)

//KafkaConsumerConfig ... structure to read kafka configuration settings
type KafkaConsumerConfig struct {
	Brokers        []string
	Topic          string
	consumer       sarama.Consumer
	client         sarama.Client
}

func main() {
	config := KafkaConsumerConfig{}
	config.Brokers = []string{"YOUR_BROKER_URL"}
	config.Topic = "YOUR_TOPIC"

	consumerConfig := sarama.NewConfig()
	consumerConfig.Net.SASL.User = "<username>"
	consumerConfig.Net.SASL.Password = "<password>"
	consumerConfig.Net.SASL.Handshake = true
	consumerConfig.Net.SASL.Enable = true

	consumerConfig.Net.TLS.Enable = true
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth: 0,
	}
	consumerConfig.Net.TLS.Config = tlsConfig

	var err error
	config.client, err = sarama.NewClient(config.Brokers, consumerConfig)
	if err != nil {
		fmt.Println("Unable to create kafka client " + err.Error())
		return
	}

	config.consumer, err = sarama.NewConsumerFromClient(config.client)
	if err != nil {
		fmt.Println("Unable to create new kafka consumer", err, config.client)
		return
	}

	partitions, err := config.client.Partitions(config.Topic)

	if err != nil {
		fmt.Println("Unable to fetch partition IDs for the topic", err, config.client, config.Topic)
		return
	}

	fmt.Println("Partitions:", partitions)

	topics, err := config.client.Topics()
	if err != nil {
		fmt.Println("Unable to fetch topics", err, config.client)
		return
	}

	fmt.Println("Topics:", topics)

}
