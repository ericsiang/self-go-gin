package kafka_test

import (
	"api/util/kafka"
	"testing"
)

func TestNewProducer(t *testing.T) {
	addr := []string{"localhost:9092"}
	kafka.NewProducer(addr, true)
	jsonS := "{\"account\":\"eric\",\"password\":\"123456\"}"
	kafka.ProducerStruct.PushMessage("mytopic", true, jsonS)
	// kafka.ProducerStruct.PushMessage("mytopic2", true, jsonS)
}
