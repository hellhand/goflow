package messaging

import "github.com/streadway/amqp"

var (
	Connection *amqp.Connection
)

func BuildMessaging(conn *amqp.Connection) {
	Connection = conn
}
