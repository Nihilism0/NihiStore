package mq

import (
	"NihiStore/server/cmd/pay/config"
	"NihiStore/server/shared/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

// MyHandler 是一个消费者类型
type MyHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	notification := &model.Order{}
	err = json.Unmarshal(msg.Body, notification)
	if err != nil {
		return err
	}
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	config.DB.Model(&model.Order{}).Create(&notification)
	return nil
}

// 初始化消费者
func InitConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	consumer := &MyHandler{
		Title: "订单消息",
	}
	c.AddHandler(consumer)
	if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
		//if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
		return err
	}
	return nil
}
