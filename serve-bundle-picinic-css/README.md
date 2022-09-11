# Serve a bundle with Capsule
> 🚧 this is a work in progress

CSS framework: [Picnic CSS](https://picnicss.com/)

## Pre-requisites

> Install dependencies
```bash
cd resources
npm install
```

> Dev mode
```bash
cd resources
npm start
```

> Build
```bash
cd resources
npm run build
```

## Build the wasm file

```bash
tinygo build -o index.wasm  -target wasi ./index.go
```
> the html code is defined in the `resources` directory

## Serve the web page

```bash
MESSAGE="👋 Hello World 🌍" capsule \
-wasm=./index.wasm \
-mode=http \
-httpPort=8080
```
> open http://localhost:8080 with your browser
