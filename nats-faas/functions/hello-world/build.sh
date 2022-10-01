#!/bin/bash
tinygo build -o hello-world.wasm -scheduler=none -target wasi ./hello-world.go

ls -lh *.wasm
