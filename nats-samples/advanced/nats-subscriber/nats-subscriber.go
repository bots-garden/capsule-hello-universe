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
  
	switch natsMessage {
	case "english":
		hf.NatsReply("🇬🇧 Hello World", 10)
	case "german":
		hf.NatsReply("🇩🇪 Hallo Welt", 10)
	case "spanish":
		hf.NatsReply("🇪🇸 Hola Mundo", 10)
	case "french":
		hf.NatsReply("🇫🇷 Bonjour le monde", 10)
	default:
		hf.NatsReply("🤔 can you repeat please?", 10)
	}

}

//export OnLoad
func OnLoad() {
	hf.Log("🙂 Hello from the NATS subscriber")
	hf.Log("👂Listening on: " + hf.NatsGetSubject())
	hf.Log("👋 NATS server: " + hf.NatsGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("👋🤗 Bye! Have a nice day")
}
