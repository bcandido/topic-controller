package topic_controller

type (
	KafkaConfig struct {
		Brokers string
	}

	TopicControllerAPI interface {
		// Creates a new topic into a kafka broker
		Create(topic Topic) error

		// Updates the number of partitions of an existing topic
		// !Be aware! Topic partitions could not be decreased
		// from a existing topic
		Update(name string, partitions int) error

		// Deletes a topic from a kafka broker
		Delete(name string) error
	}

	Topic struct {
		Name              string
		Partitions        int
		ReplicationFactor int
	}
)
