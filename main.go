package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var homeserver = flag.String("homeserver", "https://matrix.org", "Matrix homeserver")
var username = flag.String("username", "", "Matrix username localpart")
var password = flag.String("password", "", "Matrix password")
var mqttBroker = flag.String("broker", "tcp://localhost:1883", "The MQTT Broker")

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
