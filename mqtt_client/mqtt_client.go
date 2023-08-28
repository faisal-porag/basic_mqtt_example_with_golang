package mqtt_client

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
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

func (c *MQTTClient) Connect() error {
	token := c.mqttClient.Connect()
	token.Wait()
	return token.Error()
}

func (c *MQTTClient) Publish(topic string, payload string) error {
	token := c.mqttClient.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

func NewServer(brokerAddress, clientID, topic string) *MQTTClient {
	opts := MQTT.NewClientOptions().AddBroker(brokerAddress)
	opts.SetClientID(clientID)

	mqttClient := MQTT.NewClient(opts)

	return &MQTTClient{
		mqttClient: mqttClient,
	}
}

func (c *MQTTClient) Subscribe(topic string, messageHandler MQTT.MessageHandler) error {
	token := c.mqttClient.Subscribe(topic, 0, messageHandler)
	token.Wait()
	return token.Error()
}

func (c *MQTTClient) StartListening(ch chan os.Signal) {
	<-ch
	c.mqttClient.Disconnect(250)
}
