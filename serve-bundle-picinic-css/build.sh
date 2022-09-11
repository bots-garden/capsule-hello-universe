#!/bin/bash
go mod tidy
tinygo build -o index.wasm  -target wasi ./index.go

ls -lh *.wasm