package topic_controller

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"time"
)

type topicController struct {
	client *kafka.AdminClient
}

func New(config KafkaConfig) TopicControllerAPI {
	adminClient, err := newClient(config)
	if err != nil {
		return nil
	}
	return topicController{client: adminClient}
}

func newClient(config KafkaConfig) (*kafka.AdminClient, error) {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": config.Brokers})
	if err != nil {
		return nil, fmt.Errorf("error creating new kafka admin client: %w", err)
	}
	return adminClient, nil
}

func parseDuration(duration string) time.Duration {
	maxDur, err := time.ParseDuration(duration)
	if err != nil {
		panic("Duration string could not be parsed")
	}
	return maxDur
}

func (c topicController) GetAllTopics() []Topic {
	maxDur := parseDuration("60s")
	results, err := c.client.GetMetadata(nil, true, int(maxDur.Milliseconds()))

	if err != nil {
		return nil
	}

	var topics []Topic
	for _, topic := range results.Topics {
		topics = append(topics, Topic{Name: topic.Topic, Partitions: len(topic.Partitions), ReplicationFactor: len(topic.Partitions[0].Replicas)})
	}
	return topics
}

func (c topicController) Create(topic Topic) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxDur := parseDuration("60s")
	results, err := c.client.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:             topic.Name,
			NumPartitions:     topic.Partitions,
			ReplicationFactor: topic.ReplicationFactor}},
		kafka.SetAdminOperationTimeout(maxDur))

	if err != nil {
		return fmt.Errorf("error creating topic %s: %w", topic.Name, err)
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrorCode(0) {
			return fmt.Errorf("error creating topic %s: %w", result.Topic, result.Error)
		}
		log.Printf("topic created: %s\n", result.Topic)
	}

	return nil
}

func (c topicController) Update(name string, partitions int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxDur := parseDuration("60s")
	results, err := c.client.CreatePartitions(ctx, []kafka.PartitionsSpecification{{
		Topic:      name,
		IncreaseTo: partitions,
	}}, kafka.SetAdminRequestTimeout(maxDur))

	if err != nil {
		return errors.New("Failed to create partition for topic " + name + " with error:" + err.Error())
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrorCode(0) {
			return errors.New("Failed to describe topic:" + result.Error.Error())
		}
	}
	return nil
}

func (c topicController) Delete(name string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxDur := parseDuration("60s")
	results, err := c.client.DeleteTopics(
		ctx,
		[]string{name},
		kafka.SetAdminOperationTimeout(maxDur))

	if err != nil {
		return fmt.Errorf("error deleting topic %s: %w", name, err)
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrorCode(0) {
			return fmt.Errorf("error deleting topic %s: %w", result.Topic, result.Error)
		}
		log.Printf("topic created: %s\n", result.Topic)
	}

	return nil
}
