#!/bin/bash
capsule \
   -wasm=./first-subscriber.wasm \
   -mode=nats \
   -natssrv=localhost:4222 \
   -subject=ping

