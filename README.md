# Topic Controller

This project attempts implement a simple kafka topic controller allowing performing some operation into a kafka cluster like create a topic, crete new topic partitions and delete a topic.

## Topic Controller interface

```go
package main
type (
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
)
```

## Usage

```go
package main

import topic_controller "github.com/bcandido/topic-controller"

func main() {
    config := topic_controller.KafkaConfig{Brokers: "my-kafka-broker:9092"}
    controller, _ := topic_controller.New(config)
    controller.Create(Topic{Name: "topic", Partitions: 1, ReplicationFactor: 1})
}
```