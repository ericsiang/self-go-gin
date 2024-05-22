package kafka_test

import (
	"api/handler"
	"api/util/kafka"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	addr := []string{"localhost:9092"}
	kafka.NewConsumer(addr)
	kafka.ConsumerStruct.Listen("mytopic", 0, handler.ListenKafkaTest)
}
