package main

import (
	"os"
)

// EnvVar represents an environment variable name.
type EnvVar string

// Database configuration.
const (
	KafkaTopic1Var        EnvVar = "KAFKA_TOPIC_1"
	KafkaTopic2Var        EnvVar = "KAFKA_TOPIC_2"
	KafkaSeedBrokerVar    EnvVar = "KAFKA_SEED_BROKER"
	KafkaConsumerGroupVar EnvVar = "KAFKA_CONSUMER_GROUP"
)

// KafkaClientConfig represents the configuration for a Kafka client.
type KafkaClientConfig struct {
	Topic1        string
	Topic2        string
	SeedBroker    string
	ConsumerGroup string
}

// NewKafkaClientConfig creates a new KafkaClientConfig.
func NewKafkaClientConfig() *KafkaClientConfig {
	return &KafkaClientConfig{
		Topic1:        GetEnv(KafkaTopic1Var),
		Topic2:        GetEnv(KafkaTopic2Var),
		SeedBroker:    GetEnv(KafkaSeedBrokerVar),
		ConsumerGroup: GetEnv(KafkaConsumerGroupVar),
	}
}

// GetEnv returns the value of an environment variable.
func GetEnv(envVar EnvVar) string {
	return os.Getenv(string(envVar))
}
