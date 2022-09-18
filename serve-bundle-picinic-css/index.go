package main

// TinyGo wasm module
import (
	_ "embed"
	"strings"

	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

var (
	//go:embed resources/dist/index.html
	html []byte
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(hf.Request) (resp hf.Response, errResp error) {
	
	// Read the environment variable
	message, err := hf.GetEnv("MESSAGE")

	// Set the response format
	headers := map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	htmlPage :=
		strings.Replace(
			string(html),
			"{{message}}",
			message,
			-1)

	return hf.Response{Body: htmlPage, Headers: headers}, err
}
