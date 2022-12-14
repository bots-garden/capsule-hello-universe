package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {

	hf.Log("π on subject: " + hf.NatsGetSubject() + ", π message" + params[0])
  
}

//export OnLoad
func OnLoad() {
	hf.Log("π Hello from NATS subscriber")
	hf.Log("πListening on: " + hf.NatsGetSubject())
	hf.Log("π NATS server: " + hf.NatsGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("ππ€ Bye! Have a nice day")
}
