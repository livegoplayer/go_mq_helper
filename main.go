package main

import (
	"encoding/json"
	"fmt"

	"github.com/livegoplayer/go_mq_helper/rabbitmq"
)

func main() {
	//
	testPublishMessage()
	testConsumeMessage()
}

func testPublishMessage() {
	data := make(map[string]interface{})

	data["key"] = "name"
	data["value"] = "Jornery"

	msg, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	//初始化，一般放在main中
	rabbitmq.InitMqChannel("amqp://guest:guest@127.0.0.1:5670/")

	rabbitmq.Publish(&rabbitmq.Message{
		Message:    msg,
		RetryTimes: 1,
		Exchange:   "log",
		RoutingKey: "log_go_user",
	})
}

func testConsumeMessage() {
	//初始化，一般放在main中
	rabbitmq.InitMqChannel("amqp://guest:guest@127.0.0.1:5670/")

	rabbitmq.StartConsumer("log.go_user", "testlog", callback)
}

func callback(msg []byte) bool {
	data := make(map[string]interface{})
	_ = json.Unmarshal(msg, &data)

	fmt.Print(string(msg))
	fmt.Print(data)

	return true
}
