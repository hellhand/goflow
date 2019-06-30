package main

import (
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"os"
	"user-flow/messaging"
	"user-flow/rest"
	"user-flow/user-repository"
)

func main() {
	brokerURL := os.Getenv("AMQP_URL")
	if brokerURL == "" {
		brokerURL = "amqp://localhost"
	}
	connection, _ := amqp.Dial(brokerURL)
	messaging.BuildMessaging(connection)
	messaging.Read()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "localhost"
	}
	user_repository.Build(databaseURL)

	router := mux.NewRouter()
	rest.UserRest(router)

	defer connection.Close()
	defer user_repository.Close()
	defer log.Fatal(http.ListenAndServe(":9000", router))

	select {}
}
