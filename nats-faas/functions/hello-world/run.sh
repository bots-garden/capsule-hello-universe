#!/bin/bash
capsule \
   -wasm=./hello-world.wasm \
   -mode=nats \
   -natssrv=localhost:4222 \
   -subject=faas
   
