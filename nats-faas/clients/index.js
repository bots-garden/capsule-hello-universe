import { connect, StringCodec } from "nats";

console.log("🤖 Connecting to the server...")
const nc = await connect({ servers: ["localhost:4222"] });

// create a codec
const sc = StringCodec();

// subject/topic
const subject = "faas";

// the client makes a request and receives a promise for a message
// by default the request times out after 1s (1000 millis) and has
// no payload.
await nc.request(subject, sc.encode("😛 I'm a Node.js NATS client"), { timeout: 1000 })
  .then((message) => {
    console.log(`🟢 Result from the wasm function: ${sc.decode(message.data)}`);
  })
  .catch((err) => {
    console.log(`🔴 Error with request: ${err.message}`);
  });

await nc.request(subject, sc.encode("😝 It's me again"), { timeout: 1000 })
  .then((message) => {
    console.log(`🟢 Result from the wasm function: ${sc.decode(message.data)}`);
  })
  .catch((err) => {
    console.log(`🔴 Error with request: ${err.message}`);
  });

await nc.close();

