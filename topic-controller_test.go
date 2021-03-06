package topic_controller

import (
	"fmt"
	"github.com/bcandido/topic-controller/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestedController(t *testing.T) TopicControllerAPI {
	config := KafkaConfig{Brokers: "0.0.0.0:19092,0.0.0.0:29092,0.0.0.0:39092"}
	c := New(config)
	assert.NotNil(t, c)
	return c
}

func TestCreateTopic(t *testing.T) {
	topicName := rand.String(10)
	c := getTestedController(t)
	err := c.Create(Topic{Name: topicName, Partitions: 1, ReplicationFactor: 1})
	assert.Nil(t, err)
}

func TestCreateTopicManyPartitions(t *testing.T) {
	topicName := rand.String(10)
	c := getTestedController(t)
	err := c.Create(Topic{Name: topicName, Partitions: 3, ReplicationFactor: 1})
	assert.Nil(t, err)
}

func TestCreateTopicMoreReplicationFactorThanPartitions(t *testing.T) {
	topicName := rand.String(10)
	c := getTestedController(t)
	err := c.Create(Topic{Name: topicName, Partitions: 1, ReplicationFactor: 2})
	assert.Nil(t, err)
}

func TestCreateTopicMoreReplicationFactorThanPartitions_Many(t *testing.T) {
	topicName := rand.String(10)
	c := getTestedController(t)
	err := c.Create(Topic{Name: topicName, Partitions: 3, ReplicationFactor: 5})
	assert.NotNil(t, err)
}

func TestCreateTopicManyReplicationFactorAndManyPartitions(t *testing.T) {
	topicName := rand.String(10)
	c := getTestedController(t)
	err := c.Create(Topic{Name: topicName, Partitions: 16, ReplicationFactor: 3})
	assert.Nil(t, err)
}

func TestFailCreatingTopicWithSameName(t *testing.T) {
	topicName := rand.String(30)
	c := getTestedController(t)

	err := c.Create(Topic{Name: topicName, Partitions: 3, ReplicationFactor: 1})
	assert.Nil(t, err)
	err = c.Create(Topic{Name: topicName, Partitions: 3, ReplicationFactor: 1})
	assert.NotNil(t, err)
}

func TestDecreaseNumberOfPartitionsForATopic(t *testing.T) {
	topicName := rand.String(30)
	c := getTestedController(t)

	topic := Topic{Name: topicName, Partitions: 3, ReplicationFactor: 2}
	assert.Nil(t, c.Create(topic))

	err := c.Update(topic.Name, 1)
	assert.NotNil(t, err)
}

func TestIncreaseNumberOfPartitionsForATopic(t *testing.T) {
	topicName := rand.String(30)
	c := getTestedController(t)

	topic := Topic{Name: topicName, Partitions: 3, ReplicationFactor: 2}
	assert.Nil(t, c.Create(topic))

	err := c.Update(topic.Name, 7)
	assert.Nil(t, err)
}

func TestCreateAndDeleteTopic(t *testing.T) {
	topicName := rand.String(30)
	c := getTestedController(t)

	topic := Topic{Name: topicName, Partitions: 3, ReplicationFactor: 2}
	assert.Nil(t, c.Create(topic))

	err := c.Delete(topic.Name)
	assert.Nil(t, err)
}

//func TestGetTopics(t *testing.T) {
//	c := getTestedController(t)
//	topic := Topic{Name: "test", Partitions: 1, ReplicationFactor: 1}
//	expectedTopic := c.Create(topic)
//	resultTopic := c.Get(expectedTopic.Name)
//	assert.NotNil(t, resultTopic)
//}

func TestGetAllTopics(t *testing.T) {
	c := getTestedController(t)
	topics := c.GetAll()
	for _, t := range topics {
		fmt.Printf("%v:%v:%v\n", t.Name, t.Partitions, t.ReplicationFactor)
	}
}


func TestConnect(t *testing.T) {
	c := getTestedController(t)
	assert.Nil(t, c.Connect())
}