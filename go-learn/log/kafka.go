package main

import (
"fmt"
"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("测试")

	client, err := sarama.NewSyncProducer([]string{"IP:9092"}, config)
	if err != nil {
		fmt.Println("进程关闭", err)
		return
	}

	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("发送消息失败", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}