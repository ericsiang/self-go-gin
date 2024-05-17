package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type producerStruct struct {
	SyncProducer  sarama.SyncProducer
	AsyncProducer sarama.AsyncProducer
}

var ProducerStruct producerStruct

func NewProducer(addr []string, sync bool) {
	client := NewClient(addr)
	// 創建一個的生產者鏈接
	if sync {
		// 同步模式
		producer, err := sarama.NewSyncProducerFromClient(client)
		if err != nil {
			panic(err)
		}
		ProducerStruct.SyncProducer = producer
	} else {
		//非同步模式
		producer, err := sarama.NewAsyncProducerFromClient(client)
		if err != nil {
			panic(err)
		}
		ProducerStruct.AsyncProducer = producer
	}
}

func (p *producerStruct) PushMessage(topic string, sync bool, msgData interface{}) error {
	jsonData, err := json.Marshal(msgData)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(jsonData),
		Key:   sarama.ByteEncoder(topic),// 控制發到固定分區
	}

	// 發送消息
	if sync {
		// 同步
		partition, offset, err := p.SyncProducer.SendMessage(msg)
		if err != nil {
			fmt.Println("SendMessage faild : ", err)
			return err
		}
		fmt.Printf("同步發送消息成功: topic=%s, partition=%d, offset=%d\n", topic, partition, offset)
	} else {
		// 非同步
		var producer = ProducerStruct.AsyncProducer
		var producerMessage *sarama.ProducerMessage
		producer.Input() <- msg
		isSuccess := true

		func(p sarama.AsyncProducer, producerMessage *sarama.ProducerMessage) {
			select {
			case producerMessage = <-p.Successes():
				fmt.Printf("非同步發送消息成功: topic=%s, partition=%d, offset=%d\n", producerMessage.Topic, producerMessage.Partition, producerMessage.Offset)

			case fail := <-p.Errors():
				isSuccess = false
				err = fail.Err
			}
		}(producer, producerMessage)

		if !isSuccess {
			return err
		}
	}

	return nil
}
