#!/bin/bash
tinygo build -o calc.wasm -scheduler=none -target wasi ./calc.go

ls -lh *.wasm

