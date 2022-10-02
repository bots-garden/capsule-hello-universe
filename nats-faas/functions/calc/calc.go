package main

import (
	/* string to json */
	"github.com/tidwall/gjson"
	/* create json string */
	"github.com/tidwall/sjson"

	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {
	natsMessage := params[0]
	hf.Log("🖐 subject: " + hf.NatsGetSubject() + ", arguments: " + natsMessage)

	operation := gjson.Get(natsMessage, "operation").String()
	operand1 := gjson.Get(natsMessage, "operand1").Float() 
	operand2 := gjson.Get(natsMessage, "operand2").Float() 

	var res float64

	switch operation {
	case "+":
		res = operand1 + operand2
	case "-":
		res = operand1 - operand2
	case "*":
		res = operand1 * operand2
	case "/":
		res = operand1 / operand2
	default:
		res = 0.0
	}

	result := `{"result": ""}`
	result, _ = sjson.Set(result, "result", res)

	hf.Log("#️⃣ result: " + result)

	hf.NatsReply(result, 10)

}

//export OnLoad
func OnLoad() {
	hf.Log("🙂 I'm the calc function")
	hf.Log("👂Listening on: " + hf.NatsGetSubject())
	hf.Log("👋 NATS server: " + hf.NatsGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("👋🤗 Bye! Have a nice day")
}
