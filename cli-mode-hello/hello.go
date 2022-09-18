package main

import (
	"time"

	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)


// main is required.
func main() {

	hf.Log("🚀 ignition...")

	plop := func() {
		for {
			hf.Log("hello 🟢🤗 " + time.Now().String())
			time.Sleep(3500 * time.Millisecond)
		}
	}

	/*
	plouf := func() {
		for {
			hf.Log("hello 🔴🤗 " + time.Now().String())
			time.Sleep(1000 * time.Millisecond)
		}
	}
	*/

	message, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log(err.Error())
	} else {
		hf.Log("MESSAGE=" + message)
	}
	plop()
	//plouf()
	//hf.Log("...")
}

func Handle(params []string) (string, error) {

	return "ret", nil
}