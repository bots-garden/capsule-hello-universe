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
  
	switch natsMessage {
	case "english":
		hf.NatsReply("π¬π§ Hello World", 10)
	case "german":
		hf.NatsReply("π©πͺ Hallo Welt", 10)
	case "spanish":
		hf.NatsReply("πͺπΈ Hola Mundo", 10)
	case "french":
		hf.NatsReply("π«π· Bonjour le monde", 10)
	default:
		hf.NatsReply("π€ can you repeat please?", 10)
	}

}

//export OnLoad
func OnLoad() {
	hf.Log("π Hello from the NATS subscriber")
	hf.Log("πListening on: " + hf.NatsGetSubject())
	hf.Log("π NATS server: " + hf.NatsGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("ππ€ Bye! Have a nice day")
}
