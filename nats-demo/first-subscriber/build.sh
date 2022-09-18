#!/bin/bash
go mod tidy
tinygo build -o first-subscriber.wasm -scheduler=none -target wasi ./first-subscriber.go

ls -lh *.wasm
