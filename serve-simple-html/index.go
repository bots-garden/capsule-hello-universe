package main

// TinyGo wasm module
import (
	_ "embed"
	"strings"

	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

var (
	//go:embed resources/index.html
	html []byte
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(req hf.Request) (resp hf.Response, errResp error) {
	message, _ := hf.GetEnv("MESSAGE")

	headers := map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}
	resp = hf.Response{
		Body: strings.Replace(string(html), "{{message}}", message, -1),
		Headers: headers,
	}

	return resp , nil
}
