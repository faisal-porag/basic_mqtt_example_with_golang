package main

import (
	"basic_mqtt_example_with_golang/mqtt_server_client"
	"basic_mqtt_example_with_golang/utils"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := mqtt_server_client.NewServer(utils.BrokerAddress, utils.ClientID)
	if err := server.ConnectServer(); err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	messageHandler := func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("received message on topic %s: %s\n", msg.Topic(), msg.Payload())
		// Handle the message as needed (e.g., store in a database, process, etc.)
	}

	if err := server.Subscribe(utils.Topic, messageHandler); err != nil {
		fmt.Println("error subscribing:", err)
		return
	}

	server.StartListening(c)

	<-c
}
