#!/bin/bash
MESSAGE="I 💙 Lit-Element" capsule \
-wasm=./index.wasm \
-mode=http \
-httpPort=8080

gp preview $(gp url 8080)