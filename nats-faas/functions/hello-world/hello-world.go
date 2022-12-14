package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {
	natsMessage := params[0]

	hf.Log("π on subject: " + hf.NatsGetSubject() + ", π message: " + natsMessage)

    hf.NatsReply("π Hello World π", 10)
  
}

//export OnLoad
func OnLoad() {
	hf.Log("πListening on: " + hf.NatsGetSubject())
	hf.Log("π NATS server: " + hf.NatsGetServer())
}