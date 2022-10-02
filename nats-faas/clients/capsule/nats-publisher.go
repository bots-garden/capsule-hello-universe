package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
    hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {

    natsSrv := "localhost:4222"
    subject := "faas"
    operation := `{"operation":"+", "operand1":40, "operand2":2}`

    result, err := hf.NatsConnectRequest(natsSrv, subject, operation, 1)
    
    return result, err
}
