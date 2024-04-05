package main

import (
	"circle/pkg/kafka"
	"context"
	"flag"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

// Kafka client options.
const (
	SeedBroker    = "localhost:9092"
	ConsumerGroup = "my-group-identifier"
)

// Kafka topics.
const (
	KAFKA_TOPIC_1 = "first"
	KAFKA_TOPIC_2 = "second"
)

// kafkaClientOpts returns Kafka client options.
func kafkaClientOpts(consumeTopic string) []kgo.Opt {
	return []kgo.Opt{
		kgo.SeedBrokers(SeedBroker),
		kgo.ConsumerGroup(ConsumerGroup),
		kgo.ConsumeTopics(consumeTopic),
		kgo.AutoCommitMarks(),
	}
}

// main function.
func main() {
	var (
		consume string
		key     string
		msg     string
	)
	// parse flags
	flag.StringVar(&consume, "c", "", "consume Kafka topic")
	flag.StringVar(&key, "k", "", "key to send to Kafka")
	flag.StringVar(&msg, "m", "", "message to send to Kafka")
	flag.Parse()

	// consume topic1 by default
	if consume == "" {
		consume = KAFKA_TOPIC_1
	}

	// create Kafka client
	c, err := kafka.NewClient(kafkaClientOpts(consume)...)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	ctx := context.Background()

	// produce message if key or msg flags are not empty
	if msg != "" || key != "" {
		err := c.ProduceMessage(ctx, &kgo.Record{
			Topic: KAFKA_TOPIC_1,
			Key:   []byte(key),
			Value: []byte(msg),
		})
		if err != nil {
			panic(err)
		}
	}

	// poll fetches from Kafka topic
	fetch, err := c.PollFetches(ctx)
	if err != nil {
		panic(err)
	}

	// consume messages from Kafka topic
	err = c.ConsumeMessage(ctx, fetch, func(p kgo.FetchTopicPartition) {
		for _, r := range p.Records {
			fmt.Printf(
				"received message: topic=%s partition=%d offset=%d key=%s value=%s\n",
				r.Topic, r.Partition, r.Offset, r.Key, r.Value,
			)
			// redirect message from first topic to second topic
			if consume == KAFKA_TOPIC_1 {
				r.Topic = KAFKA_TOPIC_2
				if err := c.ProduceMessage(ctx, r); err != nil {
					panic(err)
				}
			}
		}
	})
	if err != nil {
		panic(err)
	}
}
