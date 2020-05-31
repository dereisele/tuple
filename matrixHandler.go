package main

import (
	"encoding/json"
	"fmt"
	"github.com/matrix-org/gomatrix"
	"strings"
)

var matrixClient *gomatrix.Client

func initMatrix() {
	c, err := gomatrix.NewClient(*homeserver, "", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	matrixClient = c

	resp, err := c.Login(&gomatrix.ReqLogin{Type: "m.login.password", User: *username, Password: *password})
	if err != nil {
		fmt.Println(err)
		return
	}

	c.SetCredentials(resp.UserID, resp.AccessToken)

	syncer := c.Syncer.(*gomatrix.DefaultSyncer)
	syncer.OnEventType("m.room.message", cbMatrix)
}

func startMatrix() {
	err := matrixClient.Sync()
	if err != nil {
		fmt.Println(err)
	}
}

var cbMatrix gomatrix.OnEventListener = func(evt *gomatrix.Event) {
	fmt.Printf("[%[5]s]: <%[1]s> %[4]s (%[2]s/%[3]s)\n", evt.Sender, evt.Type, evt.ID, evt.Content, evt.RoomID)
	go sendMqttMessage(evt)
}

func sendMqttMessage(message *gomatrix.Event) {
	topicRaw := fmt.Sprintf("_tuple/client/r0/rooms/%[1]s/event/%[2]s/%[3]s", message.RoomID, message.Type, message.ID)

	re := strings.NewReplacer("$", "")

	topic := re.Replace(topicRaw)
	payload, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	p := string(payload)
	fmt.Println(p)

	if token := mqttClient.Publish(topic, 1, false, payload); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	fmt.Print("send")
}
