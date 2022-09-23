
### Start the NATS server

```bash
nats-server --jetstream
```

### Build the NATS subscriber

```bash
cd nats-subscriber
tinygo build -o nats-subscriber.wasm -scheduler=none -target wasi ./nats-subscriber.go
```

### Run the NATS subscriber

```bash
cd nats-subscriber
capsule \
   -wasm=./nats-subscriber.wasm \
   -mode=nats \
   -natssrv=localhost:4222 \
   -subject=greetings
```
> The NATS subscriber is listening to the `greetings` subject


### Build the NATS publisher

```bash
cd nats-publisher
tinygo build -o nats-publisher.wasm -scheduler=none -target wasi ./nats-publisher.go
```

### Run the NATS publisher

```bash
capsule \
   -wasm=./nats-publisher.wasm \
   -mode=cli
```
