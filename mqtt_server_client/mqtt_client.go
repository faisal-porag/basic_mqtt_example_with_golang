package mqtt_server_client

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
)

type MQTTServer struct {
	mqttServer MQTT.Client
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
