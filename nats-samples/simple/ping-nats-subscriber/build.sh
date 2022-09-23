#!/bin/bash
tinygo build -o nats-subscriber.wasm -scheduler=none -target wasi ./nats-subscriber.go

ls -lh *.wasm
