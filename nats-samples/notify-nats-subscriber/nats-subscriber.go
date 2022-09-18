package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {

	hf.Log("🖐 on subject: " + hf.NatsGetSubject() + ", 🎉 message" + params[0])
  
}

//export OnLoad
func OnLoad() {
	hf.Log("🙂 Hello from NATS subscriber")
	hf.Log("👂Listening on: " + hf.NatsGetSubject())
	hf.Log("👋 NATS server: " + hf.NatsGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("👋🤗 Bye! Have a nice day")
}
