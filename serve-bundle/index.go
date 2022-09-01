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

func Handle(bodyReq string, headersReq map[string]string) (bodyResp string, headersResp map[string]string, errResp error) {
	message, _ := hf.GetEnv("MESSAGE")

	headersResp = map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return strings.Replace(string(html), "{{message}}", message, -1), headersResp, nil
}
