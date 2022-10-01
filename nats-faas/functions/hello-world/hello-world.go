package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {
	natsMessage := params[0]

	hf.Log("🖐 on subject: " + hf.NatsGetSubject() + ", 🎉 message: " + natsMessage)

    hf.NatsReply("👋 Hello World 🌍", 10)
  
}

//export OnLoad
func OnLoad() {
	hf.Log("👂Listening on: " + hf.NatsGetSubject())
	hf.Log("👋 NATS server: " + hf.NatsGetServer())
}