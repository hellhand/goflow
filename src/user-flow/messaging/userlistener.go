package messaging

import (
	"encoding/json"
	"fmt"
	"user-flow/model"
	"user-flow/user-repository"
)

func Read() {

	go func() {
		channel, _ := Connection.Channel()
		defer channel.Close()
		durable, exclusive := false, false
		autoDelete, noWait := true, true
		q, _ := channel.QueueDeclare("test", durable, autoDelete, exclusive, noWait, nil)
		channel.QueueBind(q.Name, "user", "amq.topic", false, nil)
		autoAck, exclusive, noLocal, noWait := false, false, false, false
		messages, _ := channel.Consume(q.Name, "", autoAck, exclusive, noLocal, noWait, nil)
		multiAck := false
		for msg := range messages {
			var user model.User
			json.Unmarshal(msg.Body, &user)
			fmt.Println("Body:", string(msg.Body), "Timestamp:", msg.Timestamp)
			fmt.Println("Body:", user, "Timestamp:", msg.Timestamp)
			user_repository.Insert(&user)
			msg.Ack(multiAck)
		}
	}()
}
