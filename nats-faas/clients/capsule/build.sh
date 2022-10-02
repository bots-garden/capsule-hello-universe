#!/bin/bash
tinygo build -o nats-publisher.wasm -scheduler=none -target wasi ./nats-publisher.go

ls -lh *.wasm
