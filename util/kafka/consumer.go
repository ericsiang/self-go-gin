package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type consumerStruct struct {
	Consumer sarama.Consumer
}

var ConsumerStruct consumerStruct

func NewConsumer(addr []string) error {
	client := NewClient(addr)

	// 創建一個的消費者鏈接
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return err
	}
	defer consumer.Close()
	ConsumerStruct.Consumer = consumer
	return nil
}

func (c consumerStruct) Listen(topic string, partition int32, handler func(value []byte)) error {
	if partition == -1 {
		// 取得某 topic 下得所有分區
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			fmt.Printf("Consumer listen get partition list, err:%v\n", err)
			return err
		}
		for partition := range partitionList { // 遍歷所有的分區
			err := c.getConsumePartition(topic, int32(partition), handler)
			if err != nil {
				return err
			}
		}
	}
	// 取得 某topic 下得指定分區
	err := c.getConsumePartition(topic, int32(partition), handler)
	if err != nil {
		return err
	}
	return nil
}

// 取得 某topic 下得指定分區
func (c consumerStruct) getConsumePartition(topic string, partition int32, handler func(value []byte)) error {
	consumePartition, err := c.Consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("Consumer listen get partition, err:%v\n", err)
		return err
	}
	// 啟用一個協程，持續監聽佇列中的數據，佇列中的資料會透過 Messages 發過來
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		select {
		case msg := <-consumePartition.Messages():
			fmt.Printf("接收topic=%s, partition=%d, offset=%d, value=%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
			// if handler(msg.Value) {
			// 	break
			// }
		case <-quit:
			fmt.Println("stop listen consumer")
			break
		}
	}
}
