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
