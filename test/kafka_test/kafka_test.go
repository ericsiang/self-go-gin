package kafka_test

import (
	"api/handler"
	"api/util/kafka"
	"testing"
)

func TestNewProducer(t *testing.T) {
	addr := []string{"localhost:9092"}
	kafka.NewProducer(addr, true)
	jsonS := "{\"account\":\"eric\",\"password\":\"123456\"}"
	kafka.ProducerStruct.PushMessage("mytopic", true, jsonS)
}

func TestNewConsumer(t *testing.T) {
	addr := []string{"localhost:9092"}
	kafka.NewConsumer(addr)
	kafka.ConsumerStruct.Listen("mytopic", -1, handler.ListenKafkaTest)
}
