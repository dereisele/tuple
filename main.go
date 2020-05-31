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
