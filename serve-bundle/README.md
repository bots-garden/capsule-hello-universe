# Serve a bundle with Capsule

## Pre-requisites

> Install parcel
```bash
cd resources
npm install --save-dev parcel
```

> Dev mode
```bash
cd resources
npx parcel ./index.html
```

> Build
```bash
cd resources
npx parcel build ./index.html
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
