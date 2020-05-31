package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
)

var mqttClient mqtt.Client

func initMqtt() {
	opts := mqtt.NewClientOptions().AddBroker(*mqttBroker).SetClientID("mx-tuple")
	mqttClient = mqtt.NewClient(opts)
}

func startMqtt() {
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	mqttClient.Subscribe("_tuple/client/r0/rooms/+/send/+", 0, handleMqttMessage)
}

var handleMqttMessage mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	go sendMatrix(message.Topic(), message.Payload())
}

func sendMatrix(topic string, message []byte) {
	// _tuple/client/r0/rooms/<roomId>/send/<eventType>
	parts := strings.Split(topic, "/")
	roomId := parts[4]
	eventType := parts[6]
	var msg map[string]interface{}

	err := json.Unmarshal(message, &msg)
	if err != nil {
		fmt.Println(err)
	}

	if _, ok := msg["msgtype"]; !ok {
		fmt.Println("No msgtype, message will not be send")
		return
	}

	_, err = matrixClient.SendMessageEvent(roomId, eventType, msg)
	if err != nil {
		fmt.Println(err)
	}
}
