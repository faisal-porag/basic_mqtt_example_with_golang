package main

import (
	"basic_mqtt_example_with_golang/mqtt_client"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := mqtt_client.NewServer(mqtt_client.BrokerAddress, mqtt_client.ClientID, mqtt_client.Topic)
	if err := server.Connect(); err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	messageHandler := func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
		// Handle the message as needed (e.g., store in a database, process, etc.)
	}

	if err := server.Subscribe(mqtt_client.Topic, messageHandler); err != nil {
		fmt.Println("Error subscribing:", err)
		return
	}

	server.StartListening(c)
}
