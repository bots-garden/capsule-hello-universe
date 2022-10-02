#!/bin/bash
capsule \
   -wasm=./calc.wasm \
   -mode=nats \
   -natssrv=localhost:4222 \
   -subject=faas
   
