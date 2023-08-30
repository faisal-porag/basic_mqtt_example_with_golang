package mqtt_client

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
)

type MQTTClient struct {
	mqttClient MQTT.Client
}

type MQTTServer struct {
	mqttServer MQTT.Client
}

func NewClient(brokerAddress, clientID string) *MQTTClient {
	opts := MQTT.NewClientOptions().AddBroker(brokerAddress)
	opts.SetClientID(clientID)

	mqttClient := MQTT.NewClient(opts)

	return &MQTTClient{
		mqttClient: mqttClient,
	}
}

func (ct *MQTTClient) ConnectClient() error {
	token := ct.mqttClient.Connect()
	token.Wait()
	return token.Error()
}

func (ct *MQTTClient) Publish(topic string, payload string) error {
	token := ct.mqttClient.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

func NewServer(brokerAddress, clientID string) *MQTTServer {
	opts := MQTT.NewClientOptions().AddBroker(brokerAddress)
	opts.SetClientID(clientID)

	mqttClient := MQTT.NewClient(opts)

	return &MQTTServer{
		mqttServer: mqttClient,
	}
}

func (s *MQTTServer) ConnectServer() error {
	token := s.mqttServer.Connect()
	token.Wait()
	return token.Error()
}

func (s *MQTTServer) Subscribe(topic string, messageHandler MQTT.MessageHandler) error {
	token := s.mqttServer.Subscribe(topic, 0, messageHandler)
	token.Wait()
	return token.Error()
}

func (s *MQTTServer) StartListening(ch chan os.Signal) {
	<-ch
	s.mqttServer.Disconnect(250)
}
