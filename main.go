package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var homeserver = flag.String("homeserver", getEnv("HOMESERVER", "https://matrix.org"), "Matrix homeserver")
var username = flag.String("username", getEnv("USERNAME", "tuple"), "Matrix username localpart")
var password = flag.String("password", getEnv("PASSWORD", ""), "Matrix password")
var mqttBroker = flag.String("broker", getEnv("BROKER", "tcp://localhost:1883"), "The MQTT Broker")

func main() {
	flag.Parse()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	initMqtt()
	initMatrix()
	go startMatrix()
	go startMqtt()
	<-c
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
