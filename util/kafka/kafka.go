package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

func getConf() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0
	// 生產訊息後是否需要通知生產者
	// 同步模式會直接傳回

	//分區,新選出一個分區
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	//ACK,發送完資料需要leader和follow都確認
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 非同步模式會回到Successes和Errors通道中
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	// 配置开启自动提交 offset，这样 samara 库会定时帮我们把最新的 offset 信息提交给 kafka
	config.Consumer.Offsets.AutoCommit.Enable = true              // 开启自动 commit offset
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动 commit时间间隔
	return config
}

// 建立 kafka 用戶端，並返回客戶端連結供生產和消費使用
func NewClient(addr []string) sarama.Client {
	client, err := sarama.NewClient(addr, getConf())
	if err != nil {
		panic(err)
	}
	return client
}
