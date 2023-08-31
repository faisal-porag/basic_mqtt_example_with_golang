package main

import (
	"basic_mqtt_example_with_golang/utils"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MQTTClient struct {
	mqttClient MQTT.Client
}

func NewClient(brokerAddress, clientID string) *MQTTClient {
	opts := MQTT.NewClientOptions().AddBroker(brokerAddress)
	opts.SetClientID(clientID)

	mqttClient := MQTT.NewClient(opts)

	return &MQTTClient{
		mqttClient: mqttClient,
	}
}

func (ct *MQTTClient) ConnectTOClient() error {
	token := ct.mqttClient.Connect()
	token.Wait()
	return token.Error()
}

func (ct *MQTTClient) Publish(topic string, payload string) error {
	token := ct.mqttClient.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

func main() {
	client := NewClient(utils.BrokerAddress, utils.ClientID)
	if err := client.ConnectTOClient(); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			latitude := generateRandomLocation()
			longitude := generateRandomLocation()

			payload := fmt.Sprintf("{\"latitude\": %f, \"longitude\": %f}", latitude, longitude)

			fmt.Printf("payload: %v", payload)
			fmt.Println("")

			err := client.Publish(utils.Topic, payload)
			if err != nil {
				fmt.Println("Error publishing:", err)
			} else {
				fmt.Println("Publish successful")
			}

			fmt.Println("==============================")

			time.Sleep(10 * time.Second)
		}
	}()

	<-c
}

func generateRandomLocation() float64 {
	return 40 + (rand.Float64() * 0.1) - 0.05
}
