package main

import (
	"circle/pkg/kafka"
	"circle/pkg/logger"
	"context"
	"flag"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
)

// Kafka client options.
const (
	SeedBroker    = "localhost:9092"
	ConsumerGroup = "my-group-identifier"
	LogLevel      = zerolog.DebugLevel // logLevel is the global log level for the application.
)

// kafkaClientOpts returns Kafka client options.
func kafkaClientOpts(config *KafkaClientConfig) []kgo.Opt {
	return []kgo.Opt{
		kgo.SeedBrokers(config.SeedBroker),
		kgo.ConsumerGroup(config.ConsumerGroup),
		kgo.ConsumeTopics(config.Topic1),
		kgo.AutoCommitMarks(),
	}
}

// logReceivedMessage logs the received message.
func logReceivedMessage(r *kgo.Record) {
	log.Info().
		Str("topic", r.Topic).
		Int32("partition", r.Partition).
		Int64("offset", r.Offset).
		Bytes("key", r.Key).
		Bytes("value", r.Value).
		Msg("received message")
}

func logProducedMessage(r *kgo.Record) {
	log.Info().
		Str("topic", r.Topic).
		Bytes("value", r.Value).
		Msg("producing to")
}

// init initializes the logger.
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	logger.Setup(LogLevel)
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

	log.Info().Msg("kafka app starting...")

	// Kafka client configuration
	cfg := NewKafkaClientConfig()

	// consume topic1 by default
	if consume != "" {
		cfg.Topic1 = consume
	}

	// create Kafka client
	c, err := kafka.NewClient(kafkaClientOpts(cfg)...)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	ctx := context.Background()

	// produce message if key or msg flags are not empty
	if msg != "" || key != "" {
		rec := &kgo.Record{
			Topic: cfg.Topic1,
			Key:   []byte(key),
			Value: []byte(msg),
		}
		if err := c.ProduceMessage(ctx, rec); err != nil {
			panic(err)
		}
		logProducedMessage(rec)
	}
	log.Info().Str("topic", cfg.Topic1).Msg("consuming from")

	// poll fetches from Kafka topic
	fetch, err := c.PollFetches(ctx)
	if err != nil {
		panic(err)
	}

	// consume messages from Kafka topic
	err = c.ConsumeMessage(ctx, fetch, func(p kgo.FetchTopicPartition) {
		for _, r := range p.Records {
			logReceivedMessage(r)

			// if consuming from topic1, send the received message to topic2
			if consume != cfg.Topic1 {
				r.Topic = cfg.Topic2
				if err := c.ProduceMessage(ctx, r); err != nil {
					panic(err)
				}
				logProducedMessage(r)
			}
		}
	})
	if err != nil {
		panic(err)
	}
}
