package messaging

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"time"
	"user-flow/model"
)

func Write(user *model.User) {

	go func(user *model.User) {
		channel, _ := Connection.Channel()
		jsonUser, _ := json.Marshal(user)
		msg := amqp.Publishing{
			DeliveryMode: 1,
			Timestamp:    time.Now(),
			ContentType:  "application/json",
			Body:         jsonUser,
		}
		_ = channel.Publish("amq.topic", "user", false, false, msg)
	}(user)
}
