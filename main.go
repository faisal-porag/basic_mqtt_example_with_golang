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

// use free public MQTT broker: broker.emqx.io

const (
	brokerAddressServerSide = "tcp://broker.emqx.io:1883"
	clientIDServerSide      = "location_server"
	topicServerSide         = "rider_location"
)

func main() {
	server := mqtt_client.NewServer(brokerAddressServerSide, clientIDServerSide, topicServerSide)
	if err := server.Connect(); err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	messageHandler := func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
		// Handle the message as needed (e.g., store in a database, process, etc.)
	}

	if err := server.Subscribe(topicServerSide, messageHandler); err != nil {
		fmt.Println("Error subscribing:", err)
		return
	}

	server.StartListening(c)
}
