# Tuple

A simple, low-level matrix <--> mqtt "bridge".

[#tuple:eiselecloud.de](https://matrix.to/#/#tuple:eiselecloud.de)

## Usage

`go get gitlab.eiselecloud.de/matrix/tuple`

`tuple -homeserver https://matrix.org -username tuple -password secret -broker tcp://localhost:1883`

The rooms have to joined manually at this point via another client.

## Matrix -> MQTT

The bridge will send all Matrix Message Events to:

Topic: `_tuple/client/r0/rooms/<roomId>/event/<eventType>/<eventId>`

Content: _the received Matrix event_, like 
```json
{
  "type": "m.room.message",
  "sender": "@tuple:eiselecloud.de",
  "content": {
    "body": "foo",
    "msgtype": "m.text"
  },
  "origin_server_ts": 1590951283344,
  "unsigned": {
    "age": 165
  },
  "event_id": "$yNhWPC-6zi85SQGpkqsbZv6K_aTW8yrpkgn0y83FgiI",
  "room_id": "!jxyndfUCCLyUOLeZuI:eiselecloud.de"
}
```

## MQTT -> Matrix

Topic: `_tuple/client/r0/rooms/<roomId>/send/<eventType>`

Content: _Any Matrix Message Event Content_, like
```json
{
    "msgtype":"m.text",
    "body": "foo"
}
```

The sender will be the `tuple` user.

Note that `tuple` will happily receive its own Matrix events and publish them to MQTT.

## Use cases

* Just for fun
* Small IOT devices, like the ESP32 that don't require persistence or other advanced Matrix features

## Design

This Bridge should closely follow the Matrix Client Server Specs, with useful exceptions.

The `eventType` is included in the mqtt topic, so that IOT devices can choose to subscribe only to compatible 
(proprietary) events

This is my first golang project, so don't judge to hard on programming decisions. 

## Useful links
* [Matrix Client Server Api Specs](https://matrix.org/docs/spec/client_server/r0.6.1)
