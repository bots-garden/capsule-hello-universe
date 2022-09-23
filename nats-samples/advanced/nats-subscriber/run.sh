#!/bin/bash
capsule \
   -wasm=./nats-subscriber.wasm \
   -mode=nats \
   -natssrv=localhost:4222 \
   -subject=greetings
