package setup

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Setup() (mqtt.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("unable to load environment variables")
	}
	brokerIP := os.Getenv("MQTT_BROKER_IP")
	brokerPort := os.Getenv("MQTT_BROKER_PORT")
	clientID := os.Getenv("MQTT_CLIENT_ID")

	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s:%s", brokerIP, brokerPort))
	options.SetClientID(clientID)
	options.SetDefaultPublishHandler(defaultPublishHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}

func connectHandler(client mqtt.Client) {
	fmt.Println("Client is connected")
}

func connectionLostHandler(client mqtt.Client, err error) {
	log.Fatalf("error: client connection was lost - %s", err)
}

func defaultPublishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Message: %s received on topic: %s\n", message.Payload(), message.Topic())
}
