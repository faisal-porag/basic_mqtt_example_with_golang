package main

import (
	"basic_mqtt_example_with_golang/mqtt_client"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	brokerAddress = "tcp://broker.emqx.io:1883"
	clientID      = "rider_tracker"
	topic         = "rider_location"
)

func main() {
	client := mqtt_client.NewClient(brokerAddress, clientID)
	if err := client.Connect(); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			latitude := generateRandomLocation()
			longitude := generateRandomLocation()

			payload := fmt.Sprintf("{\"latitude\": %f, \"longitude\": %f}", latitude, longitude)
			if err := client.Publish(topic, payload); err != nil {
				fmt.Println("Error publishing:", err)
			}

			time.Sleep(10 * time.Second)
		}
	}()

	<-c
}

func generateRandomLocation() float64 {
	return 40 + (rand.Float64() * 0.1) - 0.05
}
