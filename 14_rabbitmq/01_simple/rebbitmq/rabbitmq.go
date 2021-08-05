package rebbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//  amqp://账号:密码@rabbitmq服务器地址:端口号"
const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string //queue name
	Exchange  string //交换机
	key       string //key
	MqUrl     string
}

func NewRabbitMq(queneName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queneName,
		Exchange:  exchange,
		key:       key,
		MqUrl:     MQURL,
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "create connect err")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "get channel err")

	return rabbitmq

}

//Disconnection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", err, message)
		panic(fmt.Sprintf("%s:%s", err, message))
	}
}

//create simple MQ
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMq(queueName, "", "")
}

//
func (r *RabbitMQ) PublishSimple(message string) {
	//申请队列
	//不存在自动创建，
	//队列存在，保证消息能发送到队列
	_, err := r.channel.QueueDeclare(
		//队列名称
		r.QueueName,
		//false 消息不持久化
		false,
		//消费者都离开了，是否删除消息
		false,
		//true,队列私有性
		false,
		//发送消息是否等待服务器响应
		false,
		nil,
	)
	if err != nil {
		r.failOnErr(err, "Channel.QueueDeclare err")
	}

	//send message to queue
	r.channel.Publish(
		//交换机，默认类型是direct
		r.Exchange,
		r.QueueName,
		//true，根据交换机和routekey规则找不到对应的队列，发送的消息会返回给发送者
		false,
		//true 队列没有绑定消费者，发送的消息会返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

}
