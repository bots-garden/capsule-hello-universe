package main

import (
	"errors"
	"strings"

	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
    hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {
    var errMessages []string
    var errs error
    var result string

    spAnswer, spErr := hf.NatsConnectRequest("localhost:4222", "greetings", "spanish", 1)

    if spErr != nil {
        errMessages = append(errMessages, spErr.Error())
    } else {
        hf.Log("🟢: " + spAnswer)
    }
    
    geAnswer, geErr := hf.NatsConnectRequest("localhost:4222", "greetings", "german", 1)

    if geErr != nil {
        errMessages = append(errMessages, geErr.Error())
    } else {
        hf.Log("🟢: " + geAnswer)
    }

    engAnswer, engErr := hf.NatsConnectRequest("localhost:4222", "greetings", "english", 1)

    if engErr != nil {
        errMessages = append(errMessages, engErr.Error())
    } else {
        hf.Log("🟢: " + engAnswer)
    }

    frAnswer, frErr := hf.NatsConnectRequest("localhost:4222", "greetings", "french", 1)

    if frErr != nil {
        errMessages = append(errMessages, frErr.Error())
    } else {
        hf.Log("🟢: " + frAnswer)
    }

    if len(errMessages) > 0 {
        errsReport := strings.Join(errMessages, "|")
        hf.Log("🔴: " + errsReport)
        errs = errors.New(errsReport)
        result = ""
    } else {
        errs = nil
        result = "NATS Rocks!"
    }

    return result, errs
}
